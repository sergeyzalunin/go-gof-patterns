package builder

import "testing"

var targetPerson = Person{
	City:          "Moscow",
	StreetAddress: "Moscow st. 1",
	Postcode:      "367482",

	CompanyName:  "Google",
	Position:     "Programmer",
	AnnualIncome: 120000,
}

func TestBuildPerson(t *testing.T) {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Moscow st. 1").
		In("Moscow").
		WithPostCode("367482").
		Works().
		At("Google").
		AsA("Programmer").
		Earning(120000)
	person := *pb.Build()

	if person != targetPerson {
		t.Errorf("After creating person wants to see \n'%v' got \n'%v'", targetPerson, person)
	}
}
