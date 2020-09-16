package composite

import (
	"reflect"
	"testing"
)

func TestNewNeuronLayer(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want *NeuronLayer
	}{
		{
			name: "Neyorn 1",
			args: args{101},
			want: &NeuronLayer{make([]Neuron, 101)},
		},
		{
			name: "Neyorn 2",
			args: args{23},
			want: &NeuronLayer{make([]Neuron, 23)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNeuronLayer(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNeuronLayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeuronLayer_Iter(t *testing.T) {
	makeNeurons := func(nl *NeuronLayer) []*Neuron {
		result := make([]*Neuron, len(nl.Neurons))
		for i := range nl.Neurons {
			result[i] = &nl.Neurons[i]
		}
		return result
	}

	tests := []struct {
		name string
		n    NeuronLayer
		want []*Neuron
	}{
		{
			name: "Neyorn 1",
			n:    NeuronLayer{make([]Neuron, 1)},
		},
		{
			name: "Neyorn 2",
			n:    NeuronLayer{make([]Neuron, 0)},
		},
		{
			name: "Neyorn 3",
			n:    NeuronLayer{make([]Neuron, 100)},
		},
	}

	for i := range tests {
		tests[i].want = makeNeurons(&tests[i].n)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Iter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NeuronLayer.Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeuron_Iter(t *testing.T) {
	tests := []struct {
		name string
		n    *Neuron
		want []*Neuron
	}{
		{
			name: "Neyorn 1",
			n:    &Neuron{},
		},
		{
			name: "Neyorn 2",
			n:    &Neuron{},
		},
		{
			name: "Neyorn 3",
			n:    &Neuron{},
		},
	}

	for i := range tests {
		tests[i].want = []*Neuron{tests[i].n}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Iter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Neuron.Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	neuron1, neuron2 := NewNeuron(), NewNeuron()
	layer1, layer2 := NewNeuronLayer(2), NewNeuronLayer(4)

	type args struct {
		left  NeuronInterface
		right NeuronInterface
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "Neuron",
			args: args {
				left: neuron1,
				right: neuron2,
			},
		},
		{
			name: "NeuronLayer",
			args: args {
				left: layer1,
				right: layer2,
			},
		},
		{
			name: "Neuron and NeuronLayer",
			args: args {
				left: neuron1,
				right: layer2,
			},
		},
	}

	for i := range tests {
		tests[i].want = args{
			tests[i].args.left,
			tests[i].args.right,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Connect(tt.args.left, tt.args.right)

			if got := tt.args.left.Iter(); !reflect.DeepEqual(got, tt.want.left.Iter()) {
				t.Errorf("Neuron.Iter() = %v, want %v", got, tt.want.left.Iter())
			}

			if got := tt.args.right.Iter(); !reflect.DeepEqual(got, tt.want.right.Iter()) {
				t.Errorf("Neuron.Iter() = %v, want %v", got, tt.want.right.Iter())
			}
		})
	}
}
