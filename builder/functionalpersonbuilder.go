package builder

// FunctionalPerson is struct that we want to build
type FunctionalPerson struct {
	name, position string
}

type personMod func(*FunctionalPerson)

// FunctionalPersonBuilder is builder
type FunctionalPersonBuilder struct {
	actions []personMod
}

// Called sets the name
func (b *FunctionalPersonBuilder) Called(name string) *FunctionalPersonBuilder {
	b.actions = append(b.actions, func(p *FunctionalPerson) {
		p.name = name
	})
	return b
}

// WorksAs a position
func (b *FunctionalPersonBuilder) WorksAs(position string) *FunctionalPersonBuilder {
	b.actions = append(b.actions, func(p *FunctionalPerson) {
		p.position = position
	})
	return b
}

// Build a functional person
func (b *FunctionalPersonBuilder) Build() *FunctionalPerson {
	p := FunctionalPerson{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}