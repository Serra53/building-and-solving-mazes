package maze

import (
	"fmt"
	"testing"
)

func TestNewCell(t *testing.T) {
	c := NewCell(0, 0)
	if c.Row != 0 {
		t.Errorf("Expected Row to be 0, got %d", c.Row)
	}
	if c.Col != 0 {
		t.Errorf("Expected Col to be 0, got %d", c.Col)
	}
}

func TestLink(t *testing.T) {

	cell1 := NewCell(0, 0)
	cell2 := NewCell(0, 1)

	cell1.Link(cell2)

	if len(cell1.Links) != 1 {
		t.Errorf("Expected cell1 to have 1 link, got %d", len(cell1.Links))
	}
	if len(cell2.Links) != 1 {
		t.Errorf("Expected cell2 to have 1 link, got %d", len(cell2.Links))
	}
	if !cell1.IsLinked(cell2) {
		t.Errorf("Expected cell1 to be linked to cell2")
	}
	if !cell2.IsLinked(cell1) {
		t.Errorf("Expected cell2 to be linked to cell1")
	}
}

func TestNeighbors(t *testing.T) {

	cell1 := NewCell(0, 0)
	cell2 := NewCell(0, 1)
	cell3 := NewCell(1, 0)
	cell4 := NewCell(2, 2)

	cell1.Link(cell2)
	cell1.Link(cell3)
	cell1.Link(cell4)

	fmt.Println(cell1.Links)

	cell1.AssignNeighbors()

	if len(cell1.Neighbors) != 2 {
		t.Errorf("Expected cell1 to have 2 neighbors, got %d", len(cell1.Neighbors))
	}

	_, ok := cell1.Neighbors["E"]
	if !ok {
		t.Errorf("Expected cell1 to have neighbor cell2")
	}

	_, ok = cell1.Neighbors["W"]
	if !ok {
		t.Errorf("Expected cell1 to have neighbor cell3")
	}
}
