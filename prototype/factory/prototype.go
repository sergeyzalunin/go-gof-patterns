package factory

import (
	"bytes"
	"encoding/gob"
)

var (
	mainOffice = Employee{
		"", Address{0, "123 Est St", "Dublin"},
	}
	auxOffice = Employee{
		"", Address{0, "88 Beaker St", "London"},
	}
)

// Employee is base structure to test
type Employee struct {
	Name   string
	Office Address
}

// Address of person
type Address struct {
	Suite               int
	StreetAddress, City string
}

// DeepCopySerialize makes a deep copy of person by using serialize
func (p *Employee) DeepCopySerialize() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

// newEmployee is a prototype factory
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopySerialize()
	result.Name = name
	result.Office.Suite = suite
	return result
}

// NewMainOfficeEmployee is the main office prototype factory
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

// NeAuxOfficeEmployee is the auxiliary office prototype factory
func NeAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}
