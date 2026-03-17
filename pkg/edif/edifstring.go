package edif

type String struct {
	dataType ValueType
	value    string
}

func CreateString(value string) *String {
	var newEdifString *String

	newEdifString = new(String)
	newEdifString.value = value
	newEdifString.dataType = StringType

	return newEdifString
}

func (edifString *String) Value() any {
	return edifString.value
}

func (edifString *String) DataType() ValueType {
	return edifString.dataType
}
