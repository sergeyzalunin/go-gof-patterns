package facade

// Buffer ...
type Buffer struct {
	width, height int
	buffer        []rune
}

// NewBuffer is a constructor for a Buffer struct
func NewBuffer(width int, height int) *Buffer {
	return &Buffer{
		width,
		height,
		make([]rune, width*height),
	}
}

// At gets a rune from a buffer by provided index
func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

// Viewport ...
type Viewport struct {
	buffer *Buffer
	offset int
}

// NewViewport creates a new Viewport
func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

// GetCharacterAt gets a character shifted on offset + index
func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// Console is a facade for a buffur and viewport
type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

// NewConsole is a constuctor for a console
func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

// GetCharacterAt gets a character shifted on offset + index
func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}
