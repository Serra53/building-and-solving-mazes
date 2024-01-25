package creation

import (
	"fmt"
	"maze/maze"
)

type mazeCreation interface {
	On(grid *maze.Grid)
}

func CreateMaze(mazeCreation mazeCreation, grid *maze.Grid) {
	mazeCreation.On(grid)
}

func PrintMaze(grid *maze.Grid) {
	fmt.Println(grid.ToString())
}
