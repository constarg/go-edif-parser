package edif

type List struct {
	keyword    *Keyword // Holds the keyword used before the list value.
	identifier *Identifier
	children   []ListValue // Holds the children of the currently examined list.
	dataType   ValueType
}

func CreateComponent(
	keyword *Keyword, name *Identifier, values []ListValue,
) *List {
	var newEdifList *List

	newEdifList = new(List)
	newEdifList.keyword = keyword
	newEdifList.children = values
	newEdifList.dataType = ListType

	if name != nil {
		newEdifList.identifier = name
	}

	return newEdifList
}

func (edifComponent *List) Keyword() *Keyword {
	return edifComponent.keyword
}

func (edifComponent *List) Name() *Identifier {
	return edifComponent.identifier
}

func (edifComponent *List) Value() any {
	return edifComponent.children
}

func (edifComponent *List) DataType() ValueType {
	return edifComponent.dataType
}
