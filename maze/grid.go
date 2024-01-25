package maze

import (
	"math/rand"
	"strings"
)

type Grid struct {
	Rows    int
	Columns int
	Cells   [][]*Cell
}

func NewGrid(rows, columns int) *Grid {
	g := &Grid{Rows: rows, Columns: columns}
	g.prepareGrid()
	g.configureCells()
	return g
}

func (g *Grid) EachRow() [][]*Cell {
	return g.Cells
}

// EachCell returns a slice of pointers to Cell objects representing each cell in the grid.
func (g *Grid) EachCell() []*Cell {
	var cells []*Cell
	for i := range g.Cells {
		for j := range g.Cells[i] {
			cells = append(cells, g.Cells[i][j])
		}
	}
	return cells
}

func (g *Grid) prepareGrid() {
	g.Cells = make([][]*Cell, g.Rows)
	for i := range g.Cells {
		g.Cells[i] = make([]*Cell, g.Columns)
		for j := range g.Cells[i] {
			g.Cells[i][j] = NewCell(i, j)
		}
	}
}

func (g *Grid) configureCells() {
	for i := range g.Cells { //rows
		for j := range g.Cells[i] { //columns
			cell := g.Cells[i][j]
			cell.SetNeighbor("W", g.CellAt(i, j-1))
			cell.SetNeighbor("E", g.CellAt(i, j+1))
			cell.SetNeighbor("N", g.CellAt(i-1, j))
			cell.SetNeighbor("S", g.CellAt(i+1, j))
		}
	}
}

func (g *Grid) CellAt(row, column int) *Cell {
	if row < 0 || row >= g.Rows {
		return nil
	}
	if column < 0 || column >= g.Columns {
		return nil
	}
	return g.Cells[row][column]
}

func (g *Grid) RandomCell() *Cell {
	randX := rand.Intn(g.Rows)
	randY := rand.Intn(g.Columns)
	return g.CellAt(randX, randY)
}

func (g *Grid) Size() int {
	return g.Rows * g.Columns
}

func (g *Grid) ToString() string {
	output := "+" + strings.Repeat("---+", g.Columns) + "\n"
	for row := range g.Cells {
		top := "|"
		bottom := "+"

		for col := range g.Cells[row] {
			cell := g.Cells[row][col]
			body := "   "
			eastBoundary := "|"
			if cell.IsLinked(cell.Neighbors["E"]) {
				eastBoundary = " "
			}
			top += body + eastBoundary

			southBoundary := "---"
			if cell.IsLinked(cell.Neighbors["S"]) {
				southBoundary = "   "
			}
			corner := "+"
			bottom += southBoundary + corner
		}

		output += top + "\n"
		output += bottom + "\n"
	}
	return output
}
