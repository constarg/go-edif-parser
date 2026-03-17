package edif

import (
	"container/list"
	"fmt"
	"strconv"
)

func parse(edifRawContent []byte) (*Edif, error) {
	var (
		currEdifComponent  *Component
		completedComponent *Component

		examinedChunk     []byte
		examinedChunkType ValueType

		edifComponentStack list.List
		edifNumberValue    int

		err error
	)
	examinedChunk = nil
	currEdifComponent = new(Component)

	examinedChunkType = UnknownType

	for _, character := range edifRawContent {
		if examinedChunkType == StringType && character == ' ' {
			examinedChunk = append(examinedChunk, character)
			continue
		}

		if character == '(' || character == ')' || character == ' ' || character == '\n' {
			if examinedChunkType == IntegerType && len(examinedChunk) > 0 {
				edifNumberValue, err = strconv.Atoi(string(examinedChunk))
				if err != nil {
					return nil, err // TODO: Create a nice message.
				}

				currEdifComponent.values = append(
					currEdifComponent.values,
					CreateInteger(edifNumberValue),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == StringType && len(examinedChunk) > 0 {
				currEdifComponent.values = append(
					currEdifComponent.values,
					CreateString(string(examinedChunk)),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == KeywordType && len(examinedChunk) > 0 {
				currEdifComponent.keyword = CreateKeyword(
					string(examinedChunk),
				)

				examinedChunk = nil
				examinedChunkType = UnknownType
			} else if examinedChunkType == ComponentNameType && len(examinedChunk) > 0 {
				currEdifComponent.name = CreateName(
					string(examinedChunk),
				)

				examinedChunk = nil
				examinedChunkType = UnknownType
			}

			if character == '(' {
				if currEdifComponent != nil {
					edifComponentStack.PushBack(currEdifComponent)
				}
				examinedChunkType = KeywordType
				examinedChunk = []byte{}

				currEdifComponent = new(Component)
				currEdifComponent.values = []Value{}
				continue

			} else if character == ')' {
				completedComponent = currEdifComponent
				currEdifComponent = edifComponentStack.Back().Value.(*Component)
				currEdifComponent.values = append(
					currEdifComponent.values,
					completedComponent,
				)

				edifComponentStack.Remove(edifComponentStack.Back())
				continue
			}
			continue
		}

		if character == '"' {
			if examinedChunkType == StringType {
				fmt.Printf("%s\n", string(examinedChunk))
				currEdifComponent.values = append(
					currEdifComponent.values,
					CreateString(string(examinedChunk)),
				)
				examinedChunk = nil
				examinedChunkType = UnknownType
			} else {
				examinedChunkType = StringType
				examinedChunk = []byte{}
			}
			continue
		}

		if examinedChunkType == StringType {
			examinedChunk = append(examinedChunk, character)
		} else {
			if examinedChunk == nil {
				if character >= '0' && character <= '9' || character == '-' {
					examinedChunkType = IntegerType
				} else {
					examinedChunkType = ComponentNameType
				}
			}
			examinedChunk = append(examinedChunk, character)
		}
	}

	return nil, nil
}
