// File: edifidentifier.go
//
// **********************************************************************
//
// Defines the behaviour of an EDIF identifier.
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

type Identifier struct {
	elementType ElementType
	value       string
}

func CreateIdentifier(value string) *Identifier {
	var newEdifIdentifier *Identifier

	newEdifIdentifier = new(Identifier)
	newEdifIdentifier.value = value
	newEdifIdentifier.elementType = ListIdentifierType

	return newEdifIdentifier
}

func (edifIdentifier *Identifier) Value() any {
	return edifIdentifier.value
}

func (edifIdentifier *Identifier) DataType() ElementType {
	return edifIdentifier.elementType
}
