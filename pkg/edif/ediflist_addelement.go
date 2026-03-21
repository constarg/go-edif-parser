// File: ediflist_addelement.go
//
// **********************************************************************
//
// Implements the add element function, which allows to either push or insert
// a new element in a netlist.
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
)

// InsertElement Inserts a new netlist element after the provided, new element.
// If err = nil, the element was inserted successfully, otherwise an error is
// returned indicating the cause of failure.
func (edifList *List) InsertElement(
	element ListElement, after ListElement,
) error {
	// Holds the element of the list, after which the new element should
	// be inserted.
	var markElement *list.Element

	if edifList.children.Len() == 0 {
		return errors.New(
			"edif: no elements in the list",
		)
	}

	markElement = nil
	for curr := edifList.children.Front(); curr != nil; curr = curr.Next() {
		if curr.Value == after {
			markElement = curr
			break
		}
	}

	if markElement == nil {
		return errors.New(
			"edif: cannot insert element, no element 'after'" +
				"found",
		)
	}

	edifList.children.InsertAfter(element, markElement)
	return nil
}

func (edifList *List) PushElement(element ListElement) {
	edifList.children.PushBack(element)
}
