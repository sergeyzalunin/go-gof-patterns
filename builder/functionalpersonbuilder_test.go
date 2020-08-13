package builder

import "testing"

var originalPerson = FunctionalPerson{
	name:     "Arnold",
	position: "Actor",
}

func TestFunctionalBuilder(t *testing.T) {
	builder := FunctionalPersonBuilder{}
	got := *builder.Called("Arnold").
		WorksAs("Actor").
		Build()

	if got != originalPerson {
		t.Errorf("by creating person via builder wants to see \n'%s' got \n'%s'", originalPerson, got)
	}
}
