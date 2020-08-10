package factorygenerator

// Employee ...
type Employee struct {
	Name, Position string
	AnnualIncome int
}

// NewEmployeeFactory is functional approach
func NewEmployeeFactory(position string,
	annualIncome int) func (name string) *Employee{
		return func(name string) *Employee {
			return &Employee{name, position, annualIncome}
		}
	}


// EmployeeFactory to show structural approach
type EmployeeFactory struct {
	Position string
	AnnualIncome int
}	

// Create an employee
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// NewEmployeeFactory2 create an employee factory structure
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory  {
	return &EmployeeFactory{position, annualIncome}
}