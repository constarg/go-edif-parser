package edif

import (
	"container/list"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type incompleteList struct {
	lastElement   *list.Element
	contentStates string
}

func Write(edif *Edif) error {

	var (
		// file Represents the EDIF file.
		file *os.File

		incompleteListStack list.List

		currExaminedList    list.List
		currExaminedElement ListElement

		currIncompleteList *incompleteList

		currElement *list.Element

		currListContentState string

		err error

		errorMessage string
	)

	currListContentState += "("
	// edif.root = edif.root.children.Front().Value.(*List)
	currListContentState += edif.root.Keyword().value + " "
	currListContentState += edif.root.Identifier().value + " "

	currExaminedList = edif.root.children
	currElement = currExaminedList.Front()
	for {
		for currElement == nil && incompleteListStack.Len() != 0 {
			currListContentState += ")"
			currIncompleteList = incompleteListStack.Front().Value.(*incompleteList)
			currIncompleteList.contentStates += currListContentState

			currListContentState = currIncompleteList.contentStates + "\n"
			currElement = currIncompleteList.lastElement.Next()

			incompleteListStack.Remove(incompleteListStack.Front())
		}

		// if currElement == nil && incompleteListStack.Len() == 0 {
		// 	break
		// }

		if currElement == nil {
			break
		}``

		currExaminedElement = currElement.Value.(ListElement)

		switch currExaminedElement.DataType() {
		case StringType:
			currListContentState += " "
			currListContentState += "\""
			currListContentState += currExaminedElement.Value().(string)
			currListContentState += "\""
		case IntegerType:
			currListContentState += " "
			currListContentState += strconv.FormatInt(
				int64(currExaminedElement.Value().(int)), 10,
			)
			if currElement.Next() != nil {
				currListContentState += " "
			}
		case ListType:
			currIncompleteList = new(incompleteList)
			currIncompleteList.lastElement = currElement
			currIncompleteList.contentStates = currListContentState
			incompleteListStack.PushFront(currIncompleteList)

			currListContentState = "("
			currListContentState += currElement.Value.(*List).Keyword().value

			if currElement.Value.(*List).Identifier() != nil {
				currListContentState += " "
				currListContentState += currElement.Value.(*List).Identifier().value
				currListContentState += " "
			}

			currExaminedList = currElement.Value.(*List).children
			currElement = currExaminedList.Front()
		}

		if currExaminedElement.DataType() != ListType {
			currElement = currElement.Next()
		}
	}

	if file, err = os.OpenFile(
		edif.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644,
	); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to open file %s: %s",
			edif.filePath, err.Error(),
		)
		return errors.New(errorMessage)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Write([]byte(currListContentState))
	if err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to write file %s: %s",
			edif.filePath, err.Error(),
		)
		return errors.New(errorMessage)
	}

	return nil
}
