package edif

type Component struct {
	keyword  *Keyword // Holds the name of the keyword used before the component value.
	name     *Name
	values   []Value // Holds the component value, present after the keyword.
	dataType ValueType
}

func CreateComponent(
	keyword *Keyword, name *Name, values []Value,
) *Component {
	var newEdifComponent *Component

	newEdifComponent = new(Component)
	newEdifComponent.keyword = keyword
	newEdifComponent.values = values
	newEdifComponent.dataType = ComponentType

	if name != nil {
		newEdifComponent.name = name
	}

	return newEdifComponent
}

func (edifComponent *Component) Keyword() *Keyword {
	return edifComponent.keyword
}

func (edifComponent *Component) Name() *Name {
	return edifComponent.name
}

func (edifComponent *Component) Value() any {
	return edifComponent.values
}

func (edifComponent *Component) DataType() ValueType {
	return edifComponent.dataType
}
