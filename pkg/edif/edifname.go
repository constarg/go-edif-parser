package edif

type Identifier struct {
	dataType ValueType
	value    string
}

func CreateIdentifier(value string) *Identifier {
	var newEdifIdentifier *Identifier

	newEdifIdentifier = new(Identifier)
	newEdifIdentifier.value = value
	newEdifIdentifier.dataType = ListNameType

	return newEdifIdentifier
}

func (edifName *Identifier) Value() any {
	return edifName.value
}

func (edifName *Identifier) DataType() ValueType {
	return edifName.dataType
}
