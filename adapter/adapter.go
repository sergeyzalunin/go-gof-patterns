package adapter

import "strings"

// Line contains of coordinates of particular line
type Line struct {
	X1, Y1, X2, Y2 int
}

// VectorImage has lines to represent image
type VectorImage struct {
	Lines []Line
}

// NewRectangle create the new Vector Image
func NewRectangle(width, height int) *VectorImage {
	width--
	height--
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// ↑↑↑ the struct you are given
// ↓↓↓ the interface you have

// Point ...
type Point struct {
	X, Y int
}

// RasterImage ...
type RasterImage interface {
	GetPoints() []Point
}

// DrawPoints as string
func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX++
	maxY++

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// Adapter :
type vectoorToRasterAdapter struct {
	points []Point
}

func (v vectoorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v *vectoorToRasterAdapter) getPointers(line Line) (points []Point) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			points = append(points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			points = append(points, Point{x, top})
		}
	}
	return
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// VectorToRaster converter
func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectoorToRasterAdapter{}
	for _, line := range vi.Lines {
		points := adapter.getPointers(line)
		adapter.points = append(adapter.points, points...)
	}
	return adapter
}
