package maze

import "math/rand"

func Sidewinder(grid Grid) {
	for i := 0; i < len(grid); i++ {
		run := make([]*Cell, 0)
		for j := 0; j < len(grid[0]); j++ {
			cell := grid[i][j]
			run = append(run, cell)

			at_eastern_boundary := cell.East == nil
			at_northern_boundary := cell.North == nil

			if at_eastern_boundary || (!at_northern_boundary && rand.Intn(2) == 0) {
				member := run[rand.Intn(len(run))]
				if member.North != nil {
					member.Link(member.North)
				}
				run = make([]*Cell, 0)
			} else {
				cell.Link(cell.East)
			}

		}
	}
}
