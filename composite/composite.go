package composite

import (
	"strings"
)

// GraphicObject ...
type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

// NewFigure is a constructor to build a figure
func NewFigure(figure string, color string) *GraphicObject {
	return &GraphicObject{figure, color, nil}
}
