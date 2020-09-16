package composite

// NeuronInterface ...
type NeuronInterface interface {
	Iter() []*Neuron
}

// Neuron ...
type Neuron struct {
	In, Out []*Neuron
}

// NewNeuron is a constructor. It creates a Neuron point with
// initialized In, Out empty slices
func NewNeuron() *Neuron {
	return &Neuron {
		In: []*Neuron{},
		Out: []*Neuron{},
	}
}

// ConnectTo ...
func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

// Iter ...
func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

// NeuronLayer ...
type NeuronLayer struct {
	Neurons []Neuron
}

// NewNeuronLayer ...
func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

// Iter ...
func (n NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, len(n.Neurons))
	for i := range n.Neurons {
		result[i] = &n.Neurons[i]
	}
	return result
}

// Connect makes connections betwen two neuron interfaces
func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}
