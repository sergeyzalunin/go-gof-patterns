package bridge

import (
	"fmt"
)

// Renderer is a base interface for all renderers
type Renderer interface {
	RenderCircle(radius float32)
}

// VectorRenderer allows to render a figure in the vector format
type VectorRenderer struct {
}

// RenderCircle renders the cyrcle in the vector format
func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

// RasterRenderer is an another renderer format
type RasterRenderer struct {
	Dpi int
}

// RenderCircle renders the cyrcle as raster image
func (v *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

// Circle is the representation of Bridge pattern
type Circle struct {
	renderer Renderer
	radius   float32
}

// NewCircle is a constructor for a bridge
func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

// Draw circle by using renderer
func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

// Resize changes the radius of circle
func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}
