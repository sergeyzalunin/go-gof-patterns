package deepcopying

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []Person{
	{
		Name: "Arnold",
		Addr: &Address{
			"123 London Rd",
			"London",
			"UK",
		},
		Friends: []string{"Frank", "Jax"},
	},
	{
		Name: "Donald",
		Addr: &Address{
			"123 Usa Rd",
			"Washington",
			"USA",
		},
		Friends: []string{"Bob", "Bill"},
	},
}

func TestDeepCopying(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			person := tC.DeepCopy()
			assert.NotEqualValues(t, tC, person)
		})
	}
}

func TestDeepCopyingSerialize(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Name, func(t *testing.T) {
			person := tC.DeepCopySerialize()
			assert.NotEqualValues(t, tC, person)
		})
	}
}
