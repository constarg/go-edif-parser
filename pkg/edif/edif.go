// File: edif.go
//
// **********************************************************************
//
// Defines the behaviour of an EDIF file.
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

// ElementType Represents a datatype present in the EDIF file.
type ElementType uint32

const (
	// ListType Represents a component of the EDIF file.
	ListType ElementType = 1
	// StringType Represents a string present in the EDIF file.
	StringType ElementType = 2
	// IntegerType Represents an integer present in the EDIF file.
	IntegerType ElementType = 3
	// KeywordType Represents a keyword present in the EDIF file.
	KeywordType ElementType = 4
	// ListIdentifierType Represents the identifier of a component, present in the
	// EDIF file.
	ListIdentifierType ElementType = 5
	// UnknownType indicates that the type of datatype in the EDIF file, is
	// not yet known.
	UnknownType ElementType = 6
)

// ListElement Represents ANY datatype present in the EDIF file.
type ListElement interface {
	// Value Gets the value of the datatype (string, keyword, integer, e.t.c).
	Value() any
	// DataType Gets the datatype code (ListIdentifierType, KeywordType, e.t.c).
	DataType() ElementType
}

// Edif Models the EDIF file. Is essentially holding the whole
// tree of netlists.
type Edif struct {
	filename string // Holds the identifier of the .edf file.
	filePath string // Holds the path where the .edf file is stored.
	root     *List  // A pointer to the root netlist of the .edf.
}

// RootList Returns the root netlist, which contains all the other netlists.
func (edif *Edif) RootList() *List {
	return edif.root
}

// FileName Returns the file name of the currently examined netlist file.
func (edif *Edif) FileName() string {
	return edif.filename
}

// FilePath Returns the path where the currently examined netlist file is
// stored.
func (edif *Edif) FilePath() string {
	return edif.filePath
}
