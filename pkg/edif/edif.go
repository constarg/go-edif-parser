package edif

// EdifValueType Represents the type of the EDF value. An EDIF value
// is either an EDIF component, .e.g. is delimited in () or a string value
// delimited in "".
type EdifValueType uint32

const (
	// EdifComponentValue Represents the case where the EDIF value is an
	// EDIF component, .e.g. is delimited in ().
	EdifComponentValue EdifValueType = 1
	// EdifStringValue Represents the case where the EDIF value if a
	// string, .e.g. is delimited in "".
	EdifStringValue EdifValueType = 2
)

// EdifValue Models an EDIF value. An EDIF value can either be an
// EDIF component, or a string value.
type EdifValue struct {
	ValueName      string
	StringValue    string
	ComponentValue EdifComponent
	ValueType      EdifValueType
}

// EdifComponent Models an EDIF component. An EDIF component is
// anything that is delimited in ().
type EdifComponent struct {
	Keyword string      // Holds the name of the keyword used before the component value.
	Values  []EdifValue // Holds the component value, present after the keyword.
}

// Edif Models the EDIF file. Is essentially holding the whole
// tree of components.
type Edif struct {
	Filename      string         // Holds the name of the .edf file.
	FilePath      string         // Holds the path where the .edf file is stored.
	RootComponent *EdifComponent // A pointer to the root component of the .edf.
}
