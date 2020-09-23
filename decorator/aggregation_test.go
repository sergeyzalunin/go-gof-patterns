package decorator

import (
	"reflect"
	"testing"

	"github.com/sergeyzalunin/go-gof-patterns/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewDragon(t *testing.T) {
	tests := []struct {
		name string
		want *Dragon
	}{
		{
			name: "dragon 1",
			want: &Dragon{Bird{}, Lizard{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDragon(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDragon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDragon_Age(t *testing.T) {
	tests := []struct {
		name string
		d    Dragon
		want int
	}{
		{
			name: "dragon 1",
			d:    Dragon{Bird{age: 100}, Lizard{age: 100}},
			want: 100,
		},
		{
			name: "dragon 2",
			d:    *NewDragon(),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Age(); got != tt.want {
				t.Errorf("Dragon.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDragon_SetAge(t *testing.T) {
	type args struct {
		birdAge   int
		lizardAge int
	}
	tests := []struct {
		name string
		d    *Dragon
		age  int
		want args
	}{
		{
			name: "dragon 1",
			d:    NewDragon(),
			age:  100,
			want: args{
				birdAge:   100,
				lizardAge: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.SetAge(tt.age)
			if tt.d.bird.Age() != tt.want.birdAge {
				t.Errorf("Dragon.Age() = %v, want %v", tt.d.bird.Age(), tt.want.birdAge)
			}
			if tt.d.lizard.Age() != tt.want.lizardAge {
				t.Errorf("Dragon.Age() = %v, want %v", tt.d.lizard.Age(), tt.want.lizardAge)
			}
		})
	}
}

func TestDragon_Fly(t *testing.T) {
	tests := []struct {
		name string
		d    Dragon
		want string
	}{
		{
			name: "dragon 1",
			d: func() Dragon {
				d := *NewDragon()
				d.SetAge(100)
				return d
			}(),
			want: "Flying!\n",
		},
		{
			name: "dragon 2",
			d: func() Dragon {
				d := *NewDragon()
				d.SetAge(5)
				return d
			}(),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetPrintOutput(func() {
				tt.d.Fly()
			})
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDragon_Crawl(t *testing.T) {
	tests := []struct {
		name string
		d    *Dragon
		want string
	}{
		{
			name: "crawling dragon 1",
			d: func() *Dragon {
				d := NewDragon()
				d.SetAge(100)
				return d
			}(),
			want: "",
		},
		{
			name: "crawling dragon 2",
			d: func() *Dragon {
				d := NewDragon()
				d.SetAge(5)
				return d
			}(),
			want: "Crawling!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.GetPrintOutput(func() {
				tt.d.Crawl()
			})
			assert.Equal(t, tt.want, got)
		})
	}
}
