package creation

import (
	"math/rand"
	"maze/maze"
)

type BinaryTree struct {
	maze.Grid
}

func (bt *BinaryTree) On(grid *maze.Grid) {

	cells := grid.EachCell()
	for i := 0; i < len(cells); i++ {
		neighbors := make([]*maze.Cell, 0)
		if cells[i].Neighbors["N"] != nil {
			neighborInN := cells[i].Neighbors["N"]
			neighbors = append(neighbors, neighborInN)
		}
		if cells[i].Neighbors["E"] != nil {
			neighborInE := cells[i].Neighbors["E"]
			neighbors = append(neighbors, neighborInE)
		}

		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))
			neighbor := neighbors[index]
			cells[i].Link(neighbor)
		}
	}

	bt.Grid = *grid

}
