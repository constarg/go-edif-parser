package edif

type ValueType uint32

const (
	ComponentType ValueType = 1

	StringType ValueType = 2

	IntegerType ValueType = 3

	KeywordType ValueType = 4

	ComponentNameType ValueType = 5
	
	UnknownType ValueType = 6
)

type Value interface {
	Value() any
	DataType() ValueType
}

// Edif Models the EDIF file. Is essentially holding the whole
// tree of components.
type Edif struct {
	Filename      string     // Holds the name of the .edf file.
	FilePath      string     // Holds the path where the .edf file is stored.
	RootComponent *Component // A pointer to the root component of the .edf.
}
