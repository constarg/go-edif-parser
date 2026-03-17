package edif

type Keyword struct {
	dataType ValueType
	value    string
}

func CreateKeyword(value string) *Keyword {
	var newEdifKeyword *Keyword

	newEdifKeyword = new(Keyword)
	newEdifKeyword.value = value
	newEdifKeyword.dataType = KeywordType

	return newEdifKeyword
}

func (edifKeyword *Keyword) Value() any {
	return edifKeyword.value
}

func (edifKeyword *Keyword) DataType() ValueType {
	return edifKeyword.dataType
}
