package builder

// Person is main struct we want to get
type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

// PersonBuilder is the base structure to build a Person
type PersonBuilder struct {
	person *Person
}

// NewPersonBuilder is constructor
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// Lives provides an address builder to construct person address
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

// Works provides a job builder to construct porson job
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// Build lets complete a person
func (b *PersonBuilder) Build() *Person {
	return b.person
}

// PersonAddressBuilder builds a person address
type PersonAddressBuilder struct {
	PersonBuilder
}

// At address
func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

// In city
func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

// WithPostCode sets postcode
func (b *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// PersonJobBuilder build a person Job
type PersonJobBuilder struct {
	PersonBuilder
}

// At company
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

// AsA position
func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

// Earning some money
func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}
