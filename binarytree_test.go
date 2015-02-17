package maze

import (
	"fmt"

	"testing"
)

func TestBinaryTree(t *testing.T) {
	grid := NewGrid(4, 4)

	BinaryTree(grid)
	fmt.Println(grid)
}
