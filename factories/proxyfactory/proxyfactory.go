package proxyfactory

// Employee ...
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	// Developer ...
	Developer = iota
	// Manager ...
	Manager
)

// NewEmployee is a factory of predefined employees
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 55000}
	case Manager:
		return &Employee{"", "manager", 15000}
	default:
		panic("unsupported role")
	}
}
