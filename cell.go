package maze

type Cell struct {
	Row, Column              int
	North, South, East, West *Cell

	links map[*Cell]bool
}

func NewCell() *Cell {
	return &Cell{links: make(map[*Cell]bool)}
}

func (c *Cell) Link(cell *Cell) {
	c.links[cell] = true
	cell.links[c] = true
}

func (c *Cell) Unlink(cell *Cell) {
	delete(c.links, cell)
	delete(cell.links, c)
}

func (c *Cell) Links() []*Cell {
	links := make([]*Cell, len(c.links))
	i := 0
	for k, _ := range c.links {
		links[i] = k
		i++
	}
	return links
}

func (c *Cell) Linked(cell *Cell) bool {
	_, ok := c.links[cell]
	return ok
}

func (c *Cell) Neighbors() []*Cell {
	n := make([]*Cell, 0)
	if c.North != nil {
		n = append(n, c.North)
	}
	if c.South != nil {
		n = append(n, c.South)
	}
	if c.East != nil {
		n = append(n, c.East)
	}
	if c.West != nil {
		n = append(n, c.West)
	}
	return n
}
