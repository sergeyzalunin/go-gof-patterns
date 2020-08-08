package builder

// HTMLBuilder constructs html elements
type HTMLBuilder struct {
	rootName string
	root *HTMLElement
}

// NewHTMLBuilder serves to construct a new object
func NewHTMLBuilder(rootName string) *HTMLBuilder {
	rootElement := NewHTMLElement(rootName, "")
	return &HTMLBuilder{rootName, rootElement}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

// AddChild adds a child element to existing root elements
func (b *HTMLBuilder) AddChild(name, text string) {
	e := NewHTMLElement(name, text)
	b.root.elements = append(b.root.elements, *e)
}

// AddChildFluent adds a child element to existing root elements
// The difference between AddChild this method can make calls fluently
func (b *HTMLBuilder) AddChildFluent(name, text string) *HTMLBuilder {
	b.AddChild(name, text)
	return b
}