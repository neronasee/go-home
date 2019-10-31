package image

import "fmt"

type Driver struct{}

func (rcv Driver) PrintColor(c uint8) {
	switch c {
	case 0:
		fmt.Print("-")
	case 4:
		fmt.Print(":")
	case 9:
		fmt.Print("G")
	default:
		fmt.Print(" ")
	}
}
