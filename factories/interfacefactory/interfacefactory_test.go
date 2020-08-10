package interfacefactory

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		message	string		
	}{
		{
			message: "Person name is Arnold and age 75",		
		},
	}
	for _, tC := range testCases {
		t.Run(tC.message, func(t *testing.T)  {			
			p := NewPerson("Arnold", 75)
			got := p.SayHello()

			if got != tC.message {
				t.Errorf("wants to see %s but got %s", tC.message, got)
			}
		})
	}
}