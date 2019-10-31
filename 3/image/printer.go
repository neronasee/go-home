package image

import "fmt"

type Printer struct {
	Driver
}

func (rcv Printer) Print(i Image) {
	for index, symbol := range i.points {
		if index%i.width == 0 {
			fmt.Print("\n")
		} else {
			rcv.Driver.PrintColor(symbol.color)
		}
	}

	fmt.Println("\n")
}

func (rcv Printer) Println(i Image) {
	for _, symbol := range i.points {
		rcv.Driver.PrintColor(symbol.color)
	}

	fmt.Println("\n")
}

func NewPrinter() *Printer {
	return &Printer{
		Driver{},
	}
}
