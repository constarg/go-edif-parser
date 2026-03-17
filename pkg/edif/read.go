package edif

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Read(filepath string) (*Edif, error) {
	var (
		// file Represents the EDIF file.
		file *os.File
		// fileInfo contains various information about the file, like the
		// identifier of the file and its size.
		fileInfo fs.FileInfo

		// rootComponent Is the EDIF component which is above all other
		// components.
		rootComponent *List
		// edifFile Holds the contents of the EDIF file in the memory.
		edifFile *Edif

		// edifRawContents Holds the bytes of the EDIF file, before being
		// parsed.
		edifRawContent []byte

		// Indicates a message to be returned, when an error occurred.
		errorMessage string
		// Indicates whether an error occurred while parsing the EDIF file.
		err error
	)

	if fileInfo, err = os.Stat(filepath); os.IsNotExist(err) {
		errorMessage = fmt.Sprintf("edif: file, at %s, not exits.", filepath)
		return nil, errors.New(errorMessage)
	}

	// TODO: Check from the parsed file, the header file, and determine if the
	// file is actually an edif file.

	if file, err = os.Open(filepath); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to open file %s: %s",
			filepath, err.Error(),
		)
		return nil, errors.New(errorMessage)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	edifRawContent = make([]byte, fileInfo.Size())
	if _, err = file.Read(edifRawContent); err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to read file %s: %s",
			filepath, err.Error(),
		)
		return nil, errors.New(errorMessage)
	}

	rootComponent, err = parse(edifRawContent)
	if err != nil {
		errorMessage = fmt.Sprintf(
			"edif: failed to parse file %s: %s",
			filepath, err.Error(),
		)
		return nil, errors.New(errorMessage)
	}

	edifFile = new(Edif)
	edifFile.FilePath = filepath
	edifFile.Filename = fileInfo.Name()
	edifFile.RootComponent = rootComponent

	return edifFile, nil
}
