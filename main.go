package main

import (
	"maze/creation"
	"maze/maze"
)

func main() {
	runSW()
}

func runSW() {
	newGrid := maze.NewGrid(20, 20)
	sw := creation.Sidewinder{}
	creation.CreateMaze(&sw, newGrid)
	creation.PrintMaze(newGrid)
}


