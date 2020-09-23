package decorator

import (
	"fmt"
)

// Aged interface provides a getter and setter for age field
type Aged interface {
	Age() int
	SetAge(age int)
}

// Bird ...
type Bird struct {
	age int
}

// Age is a getter of bird age
func (b Bird) Age() int { return b.age }

// SetAge is a setter of bird age
func (b *Bird) SetAge(age int) { b.age = age }

// Fly shows that a bird can fly when it has grown
func (b Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

// Lizard ...
type Lizard struct {
	age int
}

// Age is a getter of lizard age
func (l Lizard) Age() int { return l.age }

// SetAge is a setter of lizard age
func (l *Lizard) SetAge(age int) { l.age = age }

// Crawl shows a process of lizard moving
func (l Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

// Dragon ...
type Dragon struct {
	bird   Bird
	lizard Lizard
}

// Age gets dragon's age
func (d Dragon) Age() int { return d.bird.Age() }

// SetAge sets dragon's age
func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

// Fly is a process of dragon flying
func (d Dragon) Fly() {
	d.bird.Fly()
}

// Crawl shows a process of dragon moving
func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

// NewDragon constructs a new Dragon
func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}
