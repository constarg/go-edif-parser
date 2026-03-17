package edif

type Name struct {
	dataType ValueType
	value    string
}

func CreateName(value string) *Name {
	var newEdifName *Name

	newEdifName = new(Name)
	newEdifName.value = value
	newEdifName.dataType = ListNameType

	return newEdifName
}

func (edifName *Name) Value() any {
	return edifName.value
}

func (edifName *Name) DataType() ValueType {
	return edifName.dataType
}
