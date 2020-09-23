package decorator

import (
	"fmt"
	"testing"
)

func TestCircle_Render(t *testing.T) {
	tests := []struct {
		name string
		c    *Circle
		want string
	}{
		{
			name: "rendering circle",
			c:    &Circle{25},
			want: fmt.Sprintf("Circle of radius %f", 25.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Render(); got != tt.want {
				t.Errorf("Circle.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Resize(t *testing.T) {
	type args struct {
		factor float32
	}
	tests := []struct {
		name string
		c    *Circle
		args args
		want string
	}{
		{
			name: "resizing circle",
			c:    &Circle{10},
			args: args{10},
			want: fmt.Sprintf("Circle of radius %f", 100.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Resize(tt.args.factor)
			if got := tt.c.Render(); got != tt.want {
				t.Errorf("Circle.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColoredShape_Render(t *testing.T) {
	tests := []struct {
		name string
		c    *ColoredShape
		want string
	}{
		{
			name: "rendering colored shape",
			c: &ColoredShape{
				Shape: &Circle{25},
				Color: "green",
			},
			want: fmt.Sprintf("Circle of radius %f has the color green", 25.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Render(); got != tt.want {
				t.Errorf("ColoredShape.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransparentShape_Render(t *testing.T) {
	tests := []struct {
		name string
		t    *TransparentShape
		want string
	}{
		{
			name: "rendering circle",
			t: &TransparentShape{
				Shape:        &Circle{25},
				Transparency: 100,
			},
			want: fmt.Sprintf("Circle of radius %f has %f%% transparency", 25.0, 10000.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Render(); got != tt.want {
				t.Errorf("TransparentShape.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Render(t *testing.T) {
	tests := []struct {
		name string
		s    *Square
		want string
	}{
		{
			name: "resizing square",
			s:    &Square{10},
			want: fmt.Sprintf("Square side %f", 10.0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Render(); got != tt.want {
				t.Errorf("Square.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
