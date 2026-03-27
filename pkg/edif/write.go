// File: write.go
//
// **********************************************************************
//
// Implements the read method for EDIF (Electronic Design Interchange Format).
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
	"os"
	"strconv"
)

// incomplateList Represents the state of a list if there is an interruption
// during the rebuilding of its internal components. This interruptions is
// caused by the appearance of nested lists.
type incompleteList struct {
	// The element in which the rebuilding process stopped.
	lastElement *list.Element
	// The currently built contents of the EDIF file until the point of this
	// list.
	contentStates string
}

// Write Writes the contents specified in the EDIF structure back to the
// edif file from which it was previously parsed using the Read routine. If
// err = nil, the file was successfully written, otherwise an error is returned
// indicating the cause of failure.
func Write(edif *Edif) error {

	var (
		// file Represents the EDIF file.
		file *os.File

		// Holds the incomplete lists to a stack to retrieve them later
		// to complete them, after completing their nested lists.
		incompleteListStack list.List

		// Holds the currently examined edif list.
		currExaminedList list.List
		// Holds the currently examined edif element out of the list.
		currExaminedElement ListElement

		// Holds the currently incomplete list.
		currIncompleteList *incompleteList

		// Holds the current element out of the list.
		currElement *list.Element

		// Holds the string state of the edif list.
		currListContentState string

		// Indicates a message to be returned, when an error occurred.
		errorMessage string
		// Indicates whether an error occurred while parsing the EDIF file.
		err error
	)

	// Initiate the process, by inserting the root list with the associated
	// components.
	currListContentState += "("
	if edif.root.Identifier() == nil {
		edif.root = edif.root.children.Front().Value.(*List)
	}
	currListContentState += edif.root.Keyword().value + " "
	currListContentState += edif.root.Identifier().value + " "

	currExaminedList = edif.root.children
	currElement = currExaminedList.Front()
	for {
		// While nested lists are finished, pop the previous, incomplete, list
		// out of the stack to continue the rebuilding of the whole file.
		for currElement == nil && incompleteListStack.Len() != 0 {
			currListContentState += ")"
			currIncompleteList = incompleteListStack.Front().Value.(*incompleteList)
			currIncompleteList.contentStates += currListContentState

			currListContentState = currIncompleteList.contentStates + "\n"
			currElement = currIncompleteList.lastElement.Next()

			incompleteListStack.Remove(incompleteListStack.Front())
		}

		// If no more lists remained to rebuild, exit.
		if currElement == nil {
			break
		}

		currExaminedElement = currElement.Value.(ListElement)

		switch currExaminedElement.DataType() {
		case StringType:
			currListContentState += " "
			currListContentState += "\""
			currListContentState += currExaminedElement.Value().(string)
			currListContentState += "\""
		case IntegerType:
			currListContentState += " "
			currListContentState += strconv.FormatInt(
				int64(currExaminedElement.Value().(int)), 10,
			)
			if currElement.Next() != nil {
				currListContentState += " "
			}
		case ListType:
			// When an nested list is detected, the previous, incomplete list
			// must be pushed on the front of the stack, so after the nested
			// list is completed, to continue the rebuilding from the parent
			// list.
			currIncompleteList = new(incompleteList)
			currIncompleteList.lastElement = currElement
			currIncompleteList.contentStates = currListContentState
			incompleteListStack.PushFront(currIncompleteList)

			currListContentState = "("
			currListContentState += currElement.Value.(*List).Keyword().value

			if currElement.Value.(*List).Identifier() != nil {
				currListContentState += " "
				currListContentState += currElement.Value.(*List).Identifier().value
				currListContentState += " "
			}

			currExaminedList = currElement.Value.(*List).children
			currElement = currExaminedList.Front()
		}

		if currExaminedElement.DataType() != ListType {
			currElement = currElement.Next()
		}
	}
	currListContentState += ")"

	if file, err = os.OpenFile(
		edif.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644,
	); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to open file %s: %s",
			edif.filePath, err.Error(),
		)
		return errors.New(errorMessage)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Write([]byte(currListContentState))
	if err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to write file %s: %s",
			edif.filePath, err.Error(),
		)
		return errors.New(errorMessage)
	}

	return nil
}
