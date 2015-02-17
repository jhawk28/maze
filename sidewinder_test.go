package maze

import (
	"fmt"

	"testing"
)

func TestSidewinder(t *testing.T) {
	grid := NewGrid(4, 5)

	Sidewinder(grid)
	fmt.Println(grid)
}
