package functionfactory

import "testing"

func TestFactoryFunction(t *testing.T) {
	testCases := []struct {
		person Person
	}{
		{
			person: Person{
				name:     "Arnold",
				age:      75,
				eyeCount: 2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.person.name, func(t *testing.T) {
			got := NewPerson("Arnold", 75)

			if got != tC.person {
				t.Errorf("by creating a person via factory method wants to see \n'%v' got \n'%v'", tC.person, got)
			}
		})
	}
}
