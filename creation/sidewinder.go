package creation

import (
	"math/rand"
	"maze/maze"
)

// create sidewinder algorithm
type Sidewinder struct {
	maze.Grid
}

func createStart(grid *maze.Grid) *maze.Cell {
	return grid.CellAt(0, 0)
}

func createEnd(grid *maze.Grid) *maze.Cell {
	return grid.CellAt(grid.Rows-1, grid.Columns-1)
}

func (sw *Sidewinder) On(grid *maze.Grid) {
	for _, row := range grid.EachRow() {
		run := make([]*maze.Cell, 0)
		for _, cell := range row {
			run = append(run, cell)
			atEasternBoundary := cell.Neighbors["E"] == nil
			atNorthernBoundary := cell.Neighbors["N"] == nil
			shouldCloseOut := atEasternBoundary || (!atNorthernBoundary && rand.Intn(2) == 0)
			if shouldCloseOut {
				index := rand.Intn(len(run))
				member := run[index]
				if member.Neighbors["N"] != nil {
					member.Link(member.Neighbors["N"])
				}
				run = make([]*maze.Cell, 0)
			} else {
				cell.Link(cell.Neighbors["E"])
			}
		}
	}
	sw.Grid = *grid
}
