[![Go Reference](https://pkg.go.dev/badge/github.com/constarg/go-edif-parser@v1.0.0-rc.1/pkg/edif.svg)](https://pkg.go.dev/github.com/constarg/go-edif-parser@v1.0.0-rc.1/pkg/edif)

# Intruduction 
This library supports the parsing and editting of netlist files produced by EDA tools like Vivado. It was originally made for parsing the POST-SYNTHESIS netlist produced by Xilinx Vivado. This EDIF parser supports any EDIF version, as the parsing logic is very generic.

# Installation Requirements

## Debian-basedd
```bash
sudo apt install golang-go
```

## RHEL-based
```bash
sudo dnf install golang
```
<br>

**IMPORTANT: If Go complains about the versioning, please navigate to the official Go website (https://go.dev/doc/install) and follow the instructions to install the latest version.**<br><br>

# Example 
Let's consider the following .edf content:

```edf
(edif fibex  (edifVersion 2 0 0)
  (edifLevel 0)
  (keywordMap (keywordLevel 0))
  (status
    (written
     (timeStamp 1995 11 22 23 2 53)
     (program "EDIFWRITER" (version "v8.4_2.1"))
    )
  )
  (library FIBEX
    (edifLevel 0)
    (technology
      (numberDefinition
        (scale 1 (e 1 -6) (unit distance)))
    )
    (cell dff_4 (cellType generic)
      (view view1 (viewType netlist)
        (interface)
      )
    )
)
```
Suppose we would like to add the component **(port clock (direction INPUT))** under the **dff_4** cell interface.

```GO
func example() {
	var (
		// The EDIF representation on memory.
		edif *Edif

		// Holds the components associated with the dff4 cell.
		dff4Cell *List
		// Holds the components associated with the dff4 view.
		dff4View1 *List
		// Holds the components associated with the vie1 interface.
		view1Interface *List

		// Holds the information for the newly created port.
		portList *List
		// Holds the information about the direction of the port.
		directionList *List
		// Holds the information about the children components of direction.
		portListChildren list.List

		// Indicates whether an error occurred.
		err error
	)

	edif, err = Read("/path/example.edf")
	if err != nil {
		panic(err)
	}

	allChildren := edif.RootList().ListAllChildren()
	for _, currChild := range allChildren {
		if currChild.Identifier() != nil {
			if currChild.Identifier().Value() == "dff_4" {
				dff4Cell = currChild
				break
			}
		}
	}

	allChildren = dff4Cell.ListAllChildren()
	for _, currChild := range allChildren {
		if currChild.Identifier() != nil {
			if currChild.Identifier().Value() == "view1" {
				dff4View1 = currChild
				break
			}
		}
	}

	allChildren = dff4View1.ListAllChildren()
	for _, currChild := range allChildren {
		if currChild.Keyword().Value() == "interface" {
			view1Interface = currChild
			break
		}
	}

	// (port clock (direction INPUT))
	directionList = CreateList(
		CreateKeyword("direction"),
		CreateIdentifier("INPUT"),
		list.List{},
	)
	portListChildren.PushBack(directionList)

	portList = CreateList(
		CreateKeyword("port"),
		CreateIdentifier("clock"),
		portListChildren,
	)

	view1Interface.Children().PushBack(portList)

	err = Write(edif)
	if err != nil {
		panic(err)
	}
}
```
The result after executing the code above is as follows:

```
(edif fibex (edifVersion 2  0  0)
  (edifLevel 0)
  (keywordMap(keywordLevel 0))
  (status(
    written(timeStamp 1995  11  22  23  2  53)
    (program "EDIFWRITER"(version "v8.4_2.1")))
  )
  (library FIBEX (edifLevel 0)
    (technology (numberDefinition
      (scale 1 (e 1  -6)
      (unit distance)))
    )
    (cell dff_4 (cellType generic )
        (view view1 (viewType netlist )
          (interface
              (port clock (direction INPUT))
          )
        )
    )
  )
)
```
