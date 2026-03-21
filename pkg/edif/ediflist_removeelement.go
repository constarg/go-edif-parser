// File: ediflist_addelement.go
//
// **********************************************************************
//
// Implements the remove element function, which allows to remove an elemnent
// from a netlist.
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

func (edifList *List) RemoveElement(element *ListElement) error {
	// Holds the element to be removed.
	var elementToRemove *list.Element

	for curr := edifList.children.Front(); curr != nil; curr = curr.Next() {
		if curr.Value == element {
			elementToRemove = curr
			break
		}
	}
	if elementToRemove == nil {
		return errors.New("edif: element not found")
	}

	edifList.children.Remove(elementToRemove)
	return nil
}
