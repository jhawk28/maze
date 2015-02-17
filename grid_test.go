package maze

import "testing"

func TestGridSize(t *testing.T) {
	grid := NewGrid(6, 8)

	if grid.Size() != 48 {
		t.Fail()
	}

	grid = NewGrid(1, 3)
	if grid.Size() != 3 {
		t.Fail()
	}

	grid = NewGrid(3, 1)
	if grid.Size() != 3 {
		t.Fail()
	}
}

func TestGridRandom(t *testing.T) {
	grid := NewGrid(1, 1)
	if grid[0][0] != grid.Random() {
		t.Fail()
	}
}
