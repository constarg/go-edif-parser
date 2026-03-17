package edif

type Integer struct {
	dataType ValueType
	value    int
}

func CreateInteger(value int) *Integer {
	var newEdifInteger *Integer

	newEdifInteger = new(Integer)
	newEdifInteger.value = value
	newEdifInteger.dataType = IntegerType

	return newEdifInteger
}

func (edifInteger *Integer) Value() any {
	return edifInteger.value
}

func (edifInteger *Integer) DataType() ValueType {
	return edifInteger.dataType
}
