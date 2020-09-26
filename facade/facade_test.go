package facade

import (
	"reflect"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Buffer
	}{
		{
			name: "creating a new buffer",
			args: args{100, 200},
			want: &Buffer{
				100,
				200,
				make([]rune, 100*200),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuffer(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_At(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		b    *Buffer
		args args
		want rune
	}{
		{
			name: "getting a character at",
			b:    NewBuffer(100, 200),
			args: args{0},
			want: 0,
		},
		{
			name: "getting a character at",
			b: func() *Buffer {
				b := NewBuffer(100, 200)
				b.buffer[0] = 32
				b.buffer[1] = 75
				return b
			}(),
			args: args{1},
			want: 75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.At(tt.args.index); got != tt.want {
				t.Errorf("Buffer.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewViewport(t *testing.T) {
	b := NewBuffer(100, 200)

	type args struct {
		buffer *Buffer
	}
	tests := []struct {
		name string
		args args
		want *Viewport
	}{
		{
			name: "creating a new viewport",
			args: args{b},
			want: &Viewport{
				buffer: b,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewViewport(tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewViewport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestViewport_GetCharacterAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		v    *Viewport
		args args
		want rune
	}{
		{
			name: "get a character at 0",
			v:    NewViewport(NewBuffer(100, 200)),
			args: args{0},
			want: 0,
		},
		{
			name: "getting a character at 1",
			v: func() *Viewport {
				b := NewBuffer(100, 200)
				b.buffer[0] = 32
				b.buffer[1] = 75
				return NewViewport(b)
			}(),
			args: args{1},
			want: 75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.GetCharacterAt(tt.args.index); got != tt.want {
				t.Errorf("Viewport.GetCharacterAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConsole(t *testing.T) {
	tests := []struct {
		name string
		want *Console
	}{
		{
			name: "creating a new console",
			want: func() *Console {
				b := NewBuffer(200, 150)
				v := NewViewport(b)
				return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConsole(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConsole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsole_GetCharacterAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		c    *Console
		args args
		want rune
	}{
		{
			name: "get a character at 0 index from the console",
			c: NewConsole(),
			args: args{0},
			want: 0,
		},
		{
			name: "get a character at 1 index from the console",
			c: func() *Console {
				b := NewBuffer(100, 200)
				b.buffer[0] = 32
				b.buffer[1] = 75
				v := NewViewport(b)
				return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
			}(),
			args: args{1},
			want: 75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetCharacterAt(tt.args.index); got != tt.want {
				t.Errorf("Console.GetCharacterAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
