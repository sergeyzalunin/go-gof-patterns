package builder

import "testing"

func TestNewHTMLElement(t *testing.T) {
	const target = "<ul>\n  body\n</ul>\n"
	e := NewHTMLElement("ul", "body")
	got := e.String()

	if got != target {
		t.Errorf("After creating html element wants to see \n'%s' got \n'%s'", target, got)
	}
}