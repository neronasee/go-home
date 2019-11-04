package main

import (
	"homework/4/robot"
)

func main() {
	rbt := robot.NewRobot()
	rbt.ProcessCommand(robot.PrintCurrentPositionСommand{})

	rbt.ProcessCommand(robot.SayCommand{})

	rbt.ProcessCommand(robot.UpdCoordinateCommand{Location: robot.Location{X: 3, Y: 4}})
	rbt.ProcessCommand(robot.PrintCurrentPositionСommand{})
	rbt.ProcessCommand(robot.UpdCoordinateCommand{Location: robot.Location{X: 7, Y: 7}})
	rbt.ProcessCommand(robot.PrintCurrentPositionСommand{})

	rbt.ProcessCommand(robot.PrintHistoryCommand{})

	rbt.ProcessCommand(robot.CalculateTotalDistanceCommand{})
}
