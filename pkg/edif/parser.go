// File: parser.go
//
// **********************************************************************
//
// Implements the parser for EDIF (Electronic Design Interchange Format) files.
//
// Copyright (C) 2026  Constantinos Argyriou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Email: constarg@pm.me
// ***********************************************************************

// Package edif Contains various structures, and functions related to the
// modification and parsing of EDIF (Electronic Design Interchange Format) files.
// This can be used for example to manipulate the netlist of a circuit which
// is produced by EDA tools, like Vivado.
package edif

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
)

// parse Parses the RAW contents of an .edf file. On success, it returns the
// EDIF, otherwise an error is returned indicating the cause of failure.
func parse(edifRawContent []byte) (*List, error) {
	var (
		// currEdifList Represents the component that is currently read.
		currEdifList *List
		// completeComponent Represents the component for which the parsing
		// process is completed.
		completedList *List

		// examinedChunk Represents the current portion of the EDIF file
		// to be parsed.
		examinedChunk []byte
		// examineChunkType Represents the element type of the currently examined
		// portion of the EDIF file.
		examinedChunkType ElementType

		// edifListStack Holds the non-completed components who aught to
		// be completed. It is useful to keep track of the nested components.
		edifListStack list.List
		// edifNumberValue Holds the numeric representation of a string value,
		// noticed during the parsing.
		edifNumberValue int

		// Indicates a message to be returned, when an error occurred.
		errorMessage string
		// Indicates whether an error occurred while parsing the EDIF file.
		err error
	)
	examinedChunk = nil
	currEdifList = new(List)

	examinedChunkType = UnknownType

	for _, character := range edifRawContent {
		// The strings must allow empty spaces.
		if examinedChunkType == StringType && character == ' ' {
			examinedChunk = append(examinedChunk, character)
			continue
		}

		// There are three cases by which it is safely to assume a keyword,
		// an identifier or a list's parsing is finished. This is when
		//		1. AN open bracket is detected (which means a new
		//		   component, os the datatype which is near it, can't go any
		//	       further).
		//		2. A close bracket, which is very common near integers,
		//		   therefore, it can be used to both detect where is the last
		//		   digit of a number and to complete the parsing of a component.
		//		3. An empty space, which is also very common near integers, or
		//		   after a keyword (like net, joined, rename, e.t.c.)
		// The newline is used only to be ignored.
		if character == '(' || character == ')' || character == ' ' || character == '\n' {
			if examinedChunkType == IntegerType && len(examinedChunk) > 0 {
				// Check if the element under parsing is an integer.Then,
				// append the parsed content into the component list.
				edifNumberValue, err = strconv.Atoi(string(examinedChunk))
				if err != nil {
					errorMessage = fmt.Sprintf(
						"edif: error parsing edif component number: %s",
						err,
					)

					return nil, errors.New(errorMessage)
				}

				currEdifList.children.PushBack(
					CreateInteger(edifNumberValue),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == StringType && len(examinedChunk) > 0 {
				// Check if the element under parsing a string. Then, append
				// the parsed content into the component list.
				currEdifList.children.PushBack(
					CreateString(string(examinedChunk)),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == KeywordType && len(examinedChunk) > 0 {
				// Check if the element under parsing a Keyword. Then, append
				// the parsed content into the component list.
				currEdifList.keyword = CreateKeyword(
					string(examinedChunk),
				)

				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == ListNameType && len(examinedChunk) > 0 {
				// Check if the element under parsing an identifier. Then, append
				// the parsed content into the component list.
				currEdifList.identifier = CreateIdentifier(
					string(examinedChunk),
				)

				examinedChunk = nil
				examinedChunkType = UnknownType
			}

			if character == '(' {
				// When the open parenthesis symbol is detected, two actions
				// must be taken
				//		1. To save the incomplete component into the stack, to
				//		   continue, after the completion of the nested one.
				//		2. To create the new component and fill it with the
				//		   parsed children.
				if currEdifList != nil {
					edifListStack.PushBack(currEdifList)
				}
				// After the initiation of a component ALWAYS it follows a
				// keyword (like net, renamed, joined, e.t.c).
				examinedChunkType = KeywordType
				examinedChunk = []byte{}

				currEdifList = new(List)
				// currEdifList.children
				continue
			} else if character == ')' {
				// The close parenthesis indicates that the component finished.
				// In which case, the previous component should be fetched form
				// the stack, and also to add the completed component in the
				// list of components of the one fetched from the stack (which
				// is the one in higher order of nested components).
				completedList = currEdifList
				currEdifList = edifListStack.Back().Value.(*List)
				currEdifList.children.PushBack(
					completedList,
				)

				edifListStack.Remove(edifListStack.Back())
				continue
			}
			continue
		}

		// When a quote is detected, a string is either initiated or completed.
		if character == '"' {
			// If there was another quote detected earlier, when the string
			// has been completed.
			if examinedChunkType == StringType {
				currEdifList.children.PushBack(
					CreateString(string(examinedChunk)),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else {
				// Otherwise, the string is initiated.
				examinedChunkType = StringType
				examinedChunk = []byte{}
			}
			continue
		}

		// Detect whether the currently examined chunk is a string, an integer,
		// or a component identifier.
		if examinedChunkType == StringType {
			examinedChunk = append(examinedChunk, character)
		} else {
			if examinedChunk == nil {
				if character >= '0' && character <= '9' || character == '-' {
					examinedChunkType = IntegerType
				} else {
					examinedChunkType = ListNameType
				}
			}
			examinedChunk = append(examinedChunk, character)
		}
	}

	return currEdifList, nil
}
