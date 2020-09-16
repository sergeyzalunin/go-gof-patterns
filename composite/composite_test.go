package composite

import (
	"reflect"
	"testing"
)

func TestNewFigure(t *testing.T) {
	type args struct {
		figure string
		color  string
	}
	tests := []struct {
		name string
		args args
		want *GraphicObject
	}{
		{
			name: "Drawing circle",
			args: args{"circle", "red"},
			want: &GraphicObject{
				Name:     "circle",
				Color:    "red",
				Children: nil,
			},
		},
		{
			name: "Drawing square",
			args: args{"square", "green"},
			want: &GraphicObject{
				Name:     "square",
				Color:    "green",
				Children: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFigure(tt.args.figure, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFigure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphicObject_String(t *testing.T) {
	tests := []struct {
		name string
		g    *GraphicObject
		want string
	}{
		{
			name: "Printing circle",
			g: &GraphicObject{
				Name:     "circle",
				Color:    "red",
				Children: []GraphicObject{
					{
						Name: "ellipse",
						Color: "blue",
					},
					{
						Name: "triangle",
						Color: "yellow",
					},
				},
			},
			want: "red circle\n*blue ellipse\n*yellow triangle\n",
		},
		{
			name: "Printing square",
			g: &GraphicObject{
				Name:     "square",
				Color:    "green",
			},
			want: "green square\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.String(); got != tt.want {
				t.Errorf("GraphicObject.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
