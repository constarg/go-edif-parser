// File: edifkeyword.go
//
// **********************************************************************
//
// Defines the behaviour of an EDIF keyword.
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

type Keyword struct {
	elementType ElementType
	value       string
}

func CreateKeyword(value string) *Keyword {
	var newEdifKeyword *Keyword

	newEdifKeyword = new(Keyword)
	newEdifKeyword.value = value
	newEdifKeyword.elementType = KeywordType

	return newEdifKeyword
}

func (edifKeyword *Keyword) Value() any {
	return edifKeyword.value
}

func (edifKeyword *Keyword) DataType() ElementType {
	return edifKeyword.elementType
}
