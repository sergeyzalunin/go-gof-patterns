package proxyfactory

import "testing"

type testCase struct {
	desc               string
	name, position     string
	annualIncome, role int
	want               Employee
}

var testCases = []testCase{
	{
		desc: "Arnold developer",
		name: "Arnold",
		role: Developer,
		want: Employee{
			"Arnold",
			"developer",
			55000,
		},
	},
	{
		desc: "Genry manager",
		name: "Genry",
		role: Manager,
		want: Employee{
			"Genry",
			"manager",
			15000,
		},
	},
}

func TestNewEmployeeFactory(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			emp := *NewEmployee(tC.role)
			emp.Name = tC.name

			if emp != tC.want {
				t.Errorf("We want to see %v but got %v", tC.want, emp)
			}
		})
	}
}

func TestPanic(t *testing.T) {
	testCases := []testCase{
		{
			desc: "Panic",
			name: "Panic",
			role: -999,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("We want to see panic")
				}
			}()

			_ = *NewEmployee(tC.role)
		})
	}
}
