package bridge

import (
	"reflect"
	"testing"

	"github.com/sergeyzalunin/go-gof-patterns/helper"

	"github.com/stretchr/testify/assert"
)

func TestVectorRenderer_RenderCircle(t *testing.T) {
	type args struct {
		radius float32
	}
	tests := []struct {
		name string
		v    *VectorRenderer
		args args
		want string
	}{
		{
			name: "VectorRenderer 1",
			v:    &VectorRenderer{},
			args: args{10},
			want: "Drawing a circle of radius 10\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetPrintOutput(func() {
				tt.v.RenderCircle(tt.args.radius)
			})
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRasterRenderer_RenderCircle(t *testing.T) {
	type args struct {
		radius float32
	}
	tests := []struct {
		name string
		v    *RasterRenderer
		args args
		want string
	}{
		{
			name: "RasterRenderer 1",
			v:    &RasterRenderer{},
			args: args{80},
			want: "Drawing pixels for circle of radius 80\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetPrintOutput(func() {
				tt.v.RenderCircle(tt.args.radius)
			})
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewCircle(t *testing.T) {
	type args struct {
		renderer Renderer
		radius   float32
	}
	tests := []struct {
		name string
		args args
		want *Circle
	}{
		{
			name: "Creation of new circle",
			args: args{
				renderer: &RasterRenderer{1},
				radius:   80,
			},
			want: &Circle{
				&RasterRenderer{1},
				80,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCircle(tt.args.renderer, tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Draw(t *testing.T) {
	tests := []struct {
		name string
		c    *Circle
		want string
	}{
		{
			name: "Drawing a circle of radius 10",
			c:    NewCircle(&VectorRenderer{}, 10),
			want: "Drawing a circle of radius 10\n",
		},
		{
			name: "Drawing pixels for circle of radius 80",
			c:    NewCircle(&RasterRenderer{1}, 80),
			want: "Drawing pixels for circle of radius 80\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetPrintOutput(func() {
				tt.c.Draw()
			})
			assert.Equal(t, tt.want, got)
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
		want *Circle
	}{
		{
			name: "Resizing a circle of radius 10",
			c:    NewCircle(&VectorRenderer{}, 10),
			args: args{15},
			want: NewCircle(&VectorRenderer{}, 15),
		},
		{
			name: "Resizing pixels for circle of radius 80",
			c:    NewCircle(&RasterRenderer{1}, 80),
			args: args{25},
			want: NewCircle(&RasterRenderer{1}, 25),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Resize(tt.args.factor)
			if reflect.DeepEqual(tt.c, tt.want) {
				t.Errorf("NewCircle() = %v, want %v", tt.c, tt.want)
			}
		})
	}
}
