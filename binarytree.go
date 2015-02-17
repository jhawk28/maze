package maze

import "math/rand"

func BinaryTree(grid Grid) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cell := grid[i][j]
			neighbors := make([]*Cell, 0)

			if cell.North != nil {
				neighbors = append(neighbors, cell.North)
			}
			if cell.East != nil {
				neighbors = append(neighbors, cell.East)
			}

			if len(neighbors) > 0 {
				cell.Link(neighbors[rand.Intn(len(neighbors))])
			}
		}
	}
}
