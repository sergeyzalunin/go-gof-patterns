package factorygenerator

import "testing"

type testCase struct {
	desc string
	name, position string
	annualIncome int
	want Employee
}

var testCases = []testCase {
	{
		desc: "Arnold developer",
		name: "Arnold",
		position: "developer",
		annualIncome: 55000,
		want: Employee {
			"Arnold",
			"developer",
			55000,
		},
	},
	{
		desc: "Genry manager",
		name: "Genry",
		position: "manager",
		annualIncome: 15000,
		want: Employee {
			"Genry",
			"manager",
			15000,
		},
	},
}

func TestFunctionalFactoryGenerator(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			empFactory := NewEmployeeFactory(tC.position, tC.annualIncome)
			emp := *empFactory(tC.name)

			if emp != tC.want {
				t.Errorf("We want to see %v but got %v", tC.want, emp)
			}
		})
	}
}

func TestStructuralFactoryGenerator(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			empFactory := NewEmployeeFactory2(tC.position, tC.annualIncome)
			dev := *empFactory.Create(tC.name)

			if dev != tC.want {
				t.Errorf("We want to see %v but got %v", tC.want, dev)
			}
		})
	}
}