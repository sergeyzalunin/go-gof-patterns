package builder

import (
	"strings"
	"fmt"
)

const (
	indentSize = 2
)

// HTMLElement which has base logic to display tags
type HTMLElement struct {
	name, text string
	elements []HTMLElement
}

// NewHTMLElement is the constructor for HTMLElement
func NewHTMLElement(name, text string) *HTMLElement {
	return &HTMLElement{name, text, []HTMLElement{}}
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	str := fmt.Sprintf("%s<%s>\n", i, e.name)
	sb.WriteString(str)
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent+1))
	}
	str = fmt.Sprintf("%s</%s>\n", i, e.name)
	sb.WriteString(str)

	return sb.String()
}