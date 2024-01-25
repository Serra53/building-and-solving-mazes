package maze

type Cell struct {
	Row       int
	Col       int
	Links     map[*Cell]bool
	Neighbors map[string]*Cell
}

var Directions = map[string][]int{
	"N": {-1, 0},
	"E": {0, 1},
	"S": {1, 0},
	"W": {0, -1},
}

func NewCell(row int, col int) *Cell {
	return &Cell{Row: row, Col: col, Links: make(map[*Cell]bool)}
}

func (c *Cell) Link(cell *Cell) {
	c.Links[cell] = true
	cell.Links[c] = true
}

func (c *Cell) Unlink(cell *Cell) {
	delete(c.Links, cell)
	delete(cell.Links, c)
}

func (c *Cell) ListLinks() []*Cell {
	var keys []*Cell
	for k := range c.Links {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cell) IsLinked(cell *Cell) bool {
	_, ok := c.Links[cell]
	return ok
}

func (c *Cell) SetNeighbor(dir string, cell *Cell) {
	if c.Neighbors == nil {
		c.Neighbors = make(map[string]*Cell)
	}
	c.Neighbors[dir] = cell
}

func (c *Cell) AssignNeighbors() {
	c.Neighbors = make(map[string]*Cell)
	for _, link := range c.ListLinks() {
		for dir, coord := range Directions {
			coordX, coordY := coord[0], coord[1]
			if (link.Row == c.Row+coordX) && (link.Col == c.Col+coordY) {
				c.Neighbors[dir] = link
			}
		}
	}
}
