package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	desc          string
	width, height int
	want          VectorImage
	wantRaster    []Point
	wantString    string
}

var rectangleTestCases = []testCase{
	{
		desc:   "v 1 - 6 x 4",
		width:  6,
		height: 4,
		want: VectorImage{[]Line{
			{0, 0, 5, 0},
			{0, 0, 0, 3},
			{5, 0, 5, 3},
			{0, 3, 5, 3},
		}},
		wantRaster: []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
			{X: 4, Y: 0},
			{X: 5, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 0, Y: 3},
			{X: 5, Y: 0},
			{X: 5, Y: 1},
			{X: 5, Y: 2},
			{X: 5, Y: 3},
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
			{X: 5, Y: 3},
		},
		wantString: `******
*    *
*    *
******
`,
	},
	{
		desc:   "v2 - 6 x 4",
		width:  6,
		height: 4,
		want: VectorImage{[]Line{
			{0, 0, 5, 0},
			{0, 0, 0, 3},
			{5, 0, 5, 3},
			{0, 3, 5, 3},
		}},
		wantRaster: []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
			{X: 4, Y: 0},
			{X: 5, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 0, Y: 3},
			{X: 5, Y: 0},
			{X: 5, Y: 1},
			{X: 5, Y: 2},
			{X: 5, Y: 3},
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
			{X: 5, Y: 3},
		},
		wantString: `******
*    *
*    *
******
`,
	},
	{
		desc:   "10 x 10",
		width:  10,
		height: 10,
		want: VectorImage{[]Line{
			{0, 0, 9, 0},
			{0, 0, 0, 9},
			{9, 0, 9, 9},
			{0, 9, 9, 9},
		}},
		wantRaster: []Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
			{X: 4, Y: 0},
			{X: 5, Y: 0},
			{X: 6, Y: 0},
			{X: 7, Y: 0},
			{X: 8, Y: 0},
			{X: 9, Y: 0},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 0, Y: 3},
			{X: 0, Y: 4},
			{X: 0, Y: 5},
			{X: 0, Y: 6},
			{X: 0, Y: 7},
			{X: 0, Y: 8},
			{X: 0, Y: 9},
			{X: 9, Y: 0},
			{X: 9, Y: 1},
			{X: 9, Y: 2},
			{X: 9, Y: 3},
			{X: 9, Y: 4},
			{X: 9, Y: 5},
			{X: 9, Y: 6},
			{X: 9, Y: 7},
			{X: 9, Y: 8},
			{X: 9, Y: 9},
			{X: 0, Y: 9},
			{X: 1, Y: 9},
			{X: 2, Y: 9},
			{X: 3, Y: 9},
			{X: 4, Y: 9},
			{X: 5, Y: 9},
			{X: 6, Y: 9},
			{X: 7, Y: 9},
			{X: 8, Y: 9},
			{X: 9, Y: 9},
		},
		wantString: `**********
*        *
*        *
*        *
*        *
*        *
*        *
*        *
*        *
**********
`,
	},
}

func TestCreationNewVectorImage(t *testing.T) {
	for _, tC := range rectangleTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			vectorImage := NewRectangle(tC.width, tC.height)
			assert.EqualValues(t, tC.want, *vectorImage)
		})
	}
}

func TestCreationRasterImageFromVectorImage(t *testing.T) {
	for _, tC := range rectangleTestCases {
		t.Run("Raster image: "+tC.desc, func(t *testing.T) {
			vectorImage := NewRectangle(tC.width, tC.height)
			rasterImage := VectorToRaster(vectorImage)

			rasterPoints := rasterImage.GetPoints()

			assert.EqualValues(t, tC.wantRaster, rasterPoints)
		})
	}
}

func TestCreationRasterImageFromVectorImageUsingCashe(t *testing.T) {
	for _, tC := range rectangleTestCases {
		t.Run("Raster image 1: "+tC.desc, func(t *testing.T) {
			vectorImage := NewRectangle(tC.width, tC.height)
			rasterImage := VectorToRasterCashed(vectorImage)

			rasterPoints := rasterImage.GetPoints()

			res := assert.EqualValues(t, tC.wantRaster, rasterPoints)
			t.Log(res)
		})
	}
}

func TestDrawPoints(t *testing.T) {
	for _, tC := range rectangleTestCases {
		t.Run("Raster image: "+tC.desc, func(t *testing.T) {
			vectorImage := NewRectangle(tC.width, tC.height)
			rasterImage := VectorToRaster(vectorImage)
			pointsStr := DrawPoints(rasterImage)
			assert.EqualValues(t, tC.wantString, pointsStr)
		})
	}
}
