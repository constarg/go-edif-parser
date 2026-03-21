// File: ediflist_listchildren.go
//
// **********************************************************************
//
// Implements the list children function for a netlist.
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

func (edifList *List) ListChildren() []List {
	var elements []List

	for curr := edifList.children.Front(); curr != nil; curr = curr.Next() {
		elements = append(elements, curr.Value.(List))
	}
	return elements
}
