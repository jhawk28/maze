package maze

import (
	"fmt"
	"testing"
)

func TestPng(t *testing.T) {
	grid := NewGrid(5, 5)
	BinaryTree(grid)

	fmt.Println(grid)

	m := ToImage(grid, 40)
	SaveToPngFile("hello.png", m)
}
