package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

var pointCache = map[[16]byte][]Point{}

// VectorToRasterCashed converter
func VectorToRasterCashed(vi *VectorImage) RasterImage {
	adapter := vectoorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}
	return adapter
}

func (v *vectoorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, pt := range pts {
			v.points = append(v.points, pt)
		}
		return
	}

	points := v.getPointers(line)
	v.points = append(v.points, points...)
	pointCache[h] = points
	fmt.Println("generated", len(v.points), "points")
}
