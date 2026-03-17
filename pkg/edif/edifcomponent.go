package edif

type List struct {
	tag        *Keyword // Holds the identifier of the tag used before the component value.
	identifier *Name
	children   []ListValue // Holds the component value, present after the tag.
	dataType   ValueType
}

func CreateComponent(
	keyword *Keyword, name *Name, values []ListValue,
) *List {
	var newEdifComponent *List

	newEdifComponent = new(List)
	newEdifComponent.tag = keyword
	newEdifComponent.children = values
	newEdifComponent.dataType = ListType

	if name != nil {
		newEdifComponent.identifier = name
	}

	return newEdifComponent
}

func (edifComponent *List) Keyword() *Keyword {
	return edifComponent.tag
}

func (edifComponent *List) Name() *Name {
	return edifComponent.identifier
}

func (edifComponent *List) Value() any {
	return edifComponent.children
}

func (edifComponent *List) DataType() ValueType {
	return edifComponent.dataType
}
