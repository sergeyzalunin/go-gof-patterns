package interfacefactory

import "fmt"

// Person ...
type Person interface {
	SayHello() string
}

type person struct {
	name string
	age int
}

// NewPerson ...
func NewPerson(name string, age int) Person {
	return &person{name, age}
}

func (p *person) SayHello() string {
	return fmt.Sprintf("Person name is %s and age %d", p.name, p.age)
}