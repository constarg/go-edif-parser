package edif

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Read(filepath string) (*Edif, error) {
	var (
		file     *os.File
		fileInfo fs.FileInfo

		edifRawContent []byte

		errorMessage string
		err          error
	)

	if fileInfo, err = os.Stat(filepath); os.IsNotExist(err) {
		errorMessage = fmt.Sprintf("edif: file, at %s, not exits.", filepath)
		return nil, errors.New(errorMessage)
	}

	// TODO: Check from the parsed file, the header file, and determine if the
	// file is actually an edif file.

	if file, err = os.Open(filepath); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to open file %s, %s",
			filepath, err.Error(),
		)
		return nil, errors.New(errorMessage)
	}
	defer file.Close()

	edifRawContent = make([]byte, fileInfo.Size())
	if _, err = file.Read(edifRawContent); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to read file %s, %s",
			filepath, err.Error(),
		)
		return nil, errors.New(errorMessage)
	}

	_, _ = parse(edifRawContent)

	return nil, nil
}

// for _, character := range edifRawContent {
// 	// TODO: Make sure not invalid sequence is present (like, was waiting for string and got open parenthesis.)
// 	if currEdifComponent == nil && character != '(' {
// 		continue
// 	}
//
// 	if examinedChunk == nil && examinedChunkType != UnknownType {
// 		examinedChunkType = UnknownType
// 	}
//
// 	if examinedChunkType == KeywordType && character == ' ' {
// 		currEdifComponent.keyword = CreateKeyword(string(examinedChunk))
// 		examinedChunk = nil
// 		prevExaminedChunkType = examinedChunkType
// 		continue
// 	} else if examinedChunkType == StringType && character == '"' {
// 		currEdifComponent.values = append(
// 			currEdifComponent.values,
// 			CreateString(string(examinedChunk)),
// 		)
// 		examinedChunk = nil
// 		prevExaminedChunkType = examinedChunkType
// 		continue
// 	} else if prevExaminedChunkType == KeywordType && examinedChunk != nil && (character == ' ' || character == '(' || character == '"') {
// 		fmt.Printf("%s\n", string(examinedChunk))
// 		currEdifComponent.name = CreateName(string(examinedChunk))
// 		examinedChunk = nil
// 		prevExaminedChunkType = examinedChunkType
//
// 		continue
// 	} else if examinedChunkType == IntegerType && (character == ' ' || character == ')') {
// 		edifNumberValue, err = strconv.Atoi(string(examinedChunk))
// 		if err != nil {
// 			return nil, err // TODO: Create a nice message.
// 		}
//
// 		currEdifComponent.values = append(
// 			currEdifComponent.values,
// 			CreateInteger(edifNumberValue),
// 		)
// 		examinedChunk = nil
// 		continue
// 	}
//
// 	if character == '(' {
// 		if currEdifComponent != nil {
// 			edifComponentStack.PushBack(currEdifComponent)
// 		}
// 		examinedChunkType = KeywordType
// 		examinedChunk = []byte{}
//
// 		currEdifComponent = new(Component)
// 		currEdifComponent.values = []Value{}
// 		continue
//
// 	} else if character == ')' {
// 		completedComponent = currEdifComponent
// 		currEdifComponent = edifComponentStack.Back().Value.(*Component)
// 		currEdifComponent.values = append(
// 			currEdifComponent.values,
// 			completedComponent,
// 		)
//
// 		edifComponentStack.Remove(edifComponentStack.Back())
// 		continue
// 	}
//
// 	if examinedChunkType != StringType {
// 		// Skip these characters.
// 		if character == ' ' || character == '\n' || character == '\r' {
// 			continue
// 		}
// 	}
//
// 	if character == '"' {
// 		examinedChunkType = StringType
// 		examinedChunk = []byte{}
// 		continue
// 	} else if examinedChunk == nil && character >= '0' && character <= '9' {
// 		examinedChunkType = IntegerType
// 		examinedChunk = []byte{}
// 	}
//
// 	examinedChunk = append(examinedChunk, character)
// }
