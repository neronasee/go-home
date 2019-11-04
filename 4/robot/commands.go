package robot

import (
	"fmt"
	"math"
)

type Command interface {
	Execute(r *Robot)
}

type SayCommand struct{}

func (rcv SayCommand) Execute(robot *Robot) {
	fmt.Println("Hello")
}

type UpdCoordinateCommand struct {
	Location
}

func (rcv UpdCoordinateCommand) Execute(robot *Robot) {
	fmt.Println("Moving ...")
	robot.Location = rcv.Location
	robot.MoveHistory = append(robot.MoveHistory, rcv.Location)
}

type PrintCurrentPositionCommand struct{}

func (rcv PrintCurrentPositionCommand) Execute(robot *Robot) {
	fmt.Printf("Current position: X=%d; Y=%d\n", robot.Location.X, robot.Location.Y)
}

type PrintHistoryCommand struct{}

func (rcv PrintHistoryCommand) Execute(robot *Robot) {
	fmt.Println("History:")
	for index, loc := range robot.MoveHistory {
		fmt.Printf("%d: %s\n", index, loc)
	}
}

type CalculateTotalDistanceCommand struct{}

func (rcv CalculateTotalDistanceCommand) Execute(robot *Robot) {
	var distance int

	moveHistoryWithStartingPoint := append([]Location{{0, 0}}, robot.MoveHistory...)

	for index := 0; index < len(moveHistoryWithStartingPoint)-1; index++ {
		distance = distance + rcv.calcDistance(moveHistoryWithStartingPoint[index], moveHistoryWithStartingPoint[index+1])
	}

	fmt.Printf("Total travelled distance: %d\n", distance)
}

func (rcv CalculateTotalDistanceCommand) calcDistance(start Location, end Location) int {
	return int(math.Sqrt(math.Pow(float64(end.X-start.X), 2) + math.Pow(float64(end.Y-start.Y), 2)))
}
