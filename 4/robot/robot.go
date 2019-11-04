package robot

import "fmt"

type Location struct {
	X, Y int
}

func (rcv Location) String() string {
	return fmt.Sprintf("{X: %d; Y: %d}", rcv.X, rcv.Y)
}

type Robot struct {
	Location
	MoveHistory []Location
}

func (rcv *Robot) ProcessCommand(cmd Command) {
	cmd.Execute(rcv)
}

func NewRobot() *Robot {
	return &Robot{}
}
