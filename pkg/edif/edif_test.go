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

	edif, err = Read("/home/embeddedcat/example.edf")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(edif)

	t.Skip()
}
