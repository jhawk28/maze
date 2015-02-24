package maze

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"code.google.com/p/draw2d/draw2d"
)

func ToImage(grid Grid, size float64) image.Image {
	width := int(size) * len(grid[0])
	height := int(size) * len(grid)

	i := image.NewRGBA(image.Rect(0, 0, height, width))
	gc := draw2d.NewGraphicContext(i)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cell := grid[i][j]
			x1 := float64(j) * size
			y1 := float64(i) * size
			x2 := float64(j+1) * size
			y2 := float64(i+1) * size

			if cell.North == nil {
				gc.MoveTo(x1, y1)
				gc.LineTo(x2, y1)
			}
			if cell.West == nil {
				gc.MoveTo(x1, y1)
				gc.LineTo(x1, y2)
			}
			if !cell.Linked(cell.East) {
				gc.MoveTo(x2, y1)
				gc.LineTo(x2, y2)
			}
			if !cell.Linked(cell.South) {
				gc.MoveTo(x1, y2)
				gc.LineTo(x2, y2)
			}
		}

	}
	gc.MoveTo(0, 0)
	gc.LineTo(0, float64(width))
	gc.LineTo(float64(height), float64(width))
	gc.LineTo(float64(height), 0)

	gc.Stroke()

	return i
}

func SaveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}
