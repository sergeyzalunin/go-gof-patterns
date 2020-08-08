package builder

import "testing"

const target1 = "<ul>\n</ul>\n"
const target2 = "<ul>\n  <li>\n    1\n  </li>\n  <li>\n  </li>\n</ul>\n"

func TestNewHTMLBuilder(t *testing.T) {
	b := NewHTMLBuilder("ul")
	got := b.String()

	if got != target1 {
		t.Errorf("After creating html builder wants to see \n'%s' got \n'%s'", target1, got)
	}
}

func TestAddChildElementToRoot(t *testing.T) {
	b := NewHTMLBuilder("ul")
	b.AddChild("li", "1")
	b.AddChild("li", "")

	got := b.String()

	if got != target2 {
		t.Errorf("After adding child elements wants to see \n'%s' got \n'%s'", target2, got)
	}
}

func TestAddChildElementToRootFluently(t *testing.T) {
	got := NewHTMLBuilder("ul").
		AddChildFluent("li", "1").
		AddChildFluent("li", "").
		String()

	if got != target2 {
		t.Errorf("After adding child elements wants to see \n'%s' got \n'%s'", target2, got)
	}
}