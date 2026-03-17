package edif

// ValueType Represents a datatype present in the EDIF file.
type ValueType uint32

const (
	// ComponentType Represents a component of the EDIF file.
	ComponentType ValueType = 1
	// StringType Represents a string present in the EDIF file.
	StringType ValueType = 2
	// IntegerType Represents an integer present in the EDIF file.
	IntegerType ValueType = 3
	// KeywordType Represents a keyword present in the EDIF file.
	KeywordType ValueType = 4
	// ComponentNameType Represents the name of a component, present in the
	// EDIF file.
	ComponentNameType ValueType = 5
	// UnknownType indicates that the type of datatype in the EDIF file, is
	// not yet known.
	UnknownType ValueType = 6
)

// ComponentValue Represents ANY datatype present in the EDIF file.
type ComponentValue interface {
	// Value Gets the value of the datatype (string, keyword, integer, e.t.c).
	Value() any
	// DataType Gets the datatype code (ComponentNameType, KeywordType, e.t.c).
	DataType() ValueType
}

// Edif Models the EDIF file. Is essentially holding the whole
// tree of components.
type Edif struct {
	Filename      string     // Holds the name of the .edf file.
	FilePath      string     // Holds the path where the .edf file is stored.
	RootComponent *Component // A pointer to the root component of the .edf.
}
