package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name  string
	suite int
	want  Employee
}{
	{
		"Bob",
		123,
		Employee{
			"", Address{0, "123 Est St", "Dublin"},
		},
	},
	{
		"Ted",
		88,
		Employee{
			"", Address{0, "88 Beaker St", "London"},
		},
	},
}

func TestMainOfficeProtoFactory(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := NewMainOfficeEmployee(tC.name, tC.suite)
			assert.NotEqualValues(t, tC.want, got)
		})
	}
}

func TestAuxfficeProtoFactory(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := NeAuxOfficeEmployee(tC.name, tC.suite)
			assert.NotEqualValues(t, tC.want, got)
		})
	}
}
