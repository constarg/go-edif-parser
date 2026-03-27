// File: ediflist.go
//
// **********************************************************************
//
// Defines the behaviour of an EDIF List.
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

import "container/list"

type List struct {
	keyword     *Keyword // Holds the keyword used before the list value.
	identifier  *Identifier
	children    list.List // Holds the children of the currently examined list.
	elementType ElementType
}

func CreateList(
	keyword *Keyword, identifier *Identifier,
	values list.List,
) *List {
	var newEdifList *List

	newEdifList = new(List)
	newEdifList.keyword = keyword
	newEdifList.children = values
	newEdifList.elementType = ListType

	if identifier != nil {
		newEdifList.identifier = identifier
	}

	return newEdifList
}

func (edifList *List) Children() *list.List {
	return &edifList.children
}

func (edifList *List) Keyword() *Keyword {
	return edifList.keyword
}

func (edifList *List) Identifier() *Identifier {
	return edifList.identifier
}

func (edifList *List) Value() any {
	var (
		edifElements    []ListElement
		currListElement *list.Element
	)
	currListElement = edifList.children.Front()

	for ; currListElement != nil; currListElement = currListElement.Next() {
		edifElements = append(
			edifElements, currListElement.Value.(ListElement),
		)
	}

	return edifElements
}

func (edifList *List) DataType() ElementType {
	return edifList.elementType
}
