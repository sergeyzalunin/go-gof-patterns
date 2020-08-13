package functionfactory

// Person ...
type Person struct {
	name     string
	age      int
	eyeCount int
}

// NewPerson is a factory method
func NewPerson(name string, age int) Person {
	return Person{name, age, 2}
}
