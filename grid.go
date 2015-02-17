package maze

import (
	"math/rand"
	"time"
)

type Grid [][]*Cell

func NewGrid(rows, columns int) Grid {
	grid := make(Grid, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]*Cell, columns)
		for j := 0; j < columns; j++ {
			grid[i][j] = NewCell()
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i != 0 {
				grid[i][j].North = grid[i-1][j]
			}
			if i != rows-1 {
				grid[i][j].South = grid[i+1][j]
			}
			if j != 0 {
				grid[i][j].West = grid[i][j-1]
			}
			if j != columns-1 {
				grid[i][j].East = grid[i][j+1]
			}
		}
	}

	return grid
}

func (g Grid) Random() *Cell {

	row := rand.Intn(len(g))
	col := rand.Intn(len(g[0]))

	return g[row][col]
}

func (g Grid) Size() int {
	return len(g) * len(g[0])
}

func (g Grid) String() string {
	output := "+"
	for i := 0; i < len(g[0]); i++ {
		output += "---+"
	}
	output += "\n"
	for i := 0; i < len(g); i++ {
		top := "|"
		bottom := "+"

		for j := 0; j < len(g[0]); j++ {
			cell := g[i][j]
			body := "   "
			east_boundary := "|"
			if cell.Linked(cell.East) {
				east_boundary = " "
			}
			top += body + east_boundary

			south_boundary := "---"
			if cell.Linked(cell.South) {
				south_boundary = "   "
			}
			corner := "+"
			bottom += south_boundary + corner
		}

		output += top + "\n"
		output += bottom + "\n"
	}
	return output
}

func init() {
	rand.Seed(time.Now().Unix())
}
