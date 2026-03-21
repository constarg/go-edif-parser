package edif

import (
	"fmt"
	"testing"
)

func TestEdif(t *testing.T) {
	var (
		edif *Edif
		err  error
	)

	edif, err = Read("/home/embeddedcat/Documents/Personal/University/unipi/MSc_Thesis/netlists/tester.dcp_FILES/tester.edf")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(edif)
	// var test []ListElement
	//
	// test = edif.root.Value().([]ListElement)

	// fmt.Println(len(test))
	// fmt.Println(test[0].Value().([]ListElement)[0].DataType())

	t.Skip()
}
