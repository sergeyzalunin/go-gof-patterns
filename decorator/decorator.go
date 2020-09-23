package decorator

import (
	"fmt"
	"math"
)

// Shape is the base interface to render a figure
type Shape interface {
	Render() string
}

// Circle ...
type Circle struct {
	Radius float32
}

// Render a circle
func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

// Resize a circle
func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

// Square is another one figure
type Square struct {
	Side float32
}

// Render renders a square shape
func (s *Square) Render() string {
	return fmt.Sprintf("Square side %f", s.Side)
}

// ColoredShape is the first decorator
type ColoredShape struct {
	Shape Shape
	Color string
}

//Render renders the ColoredShape interface
func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

// TransparentShape is an another type of decorator which adds a transparency to a shape
type TransparentShape struct {
	Shape Shape
	Transparency float64
}

// Render is rendering a transparent shape
func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), math.Ceil(t.Transparency * 100.0))
}