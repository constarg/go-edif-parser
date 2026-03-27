// File: edifinteger.go
//
// **********************************************************************
//
// Defines the behaviour of a EDIF integers.
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

package edif

import "container/list"

// ListAllChildren Stores every component of the EDIF file into a slice of
// pointers by doing there is no requirement to know the nested order of the
// EDIF to access a specific component. It returns the slice of pointers to
// all the components.
func (edifList *List) ListAllChildren() []*List {
	var (
		// The slice of pointers which contains all the components.
		allChildren []*List

		// The list children which should be accessed, but have not yet.
		listChildrenQueue list.List
	)

	allChildren = append(allChildren, edifList)
	listChildrenQueue.PushBack(edifList)

	currList := listChildrenQueue.Front()
	for ; currList != nil; currList = currList.Next() {
		for _, currChild := range currList.Value.(*List).ListChildren() {
			if currChild.DataType() == ListType {
				allChildren = append(allChildren, currChild.(*List))
				listChildrenQueue.PushBack(currChild)
			}
		}
	}

	return allChildren
}
