package deepcopying

// Person is base structure to test
type Person struct {
	Name    string
	Addr    *Address
	Friends []string
}

// Address of person
type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Addr = p.Addr.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}
