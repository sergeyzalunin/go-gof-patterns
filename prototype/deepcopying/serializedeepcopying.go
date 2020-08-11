package deepcopying

import (
	"bytes"
	"encoding/gob"
)

// DeepCopySerialize makes a deep copy of person by using serialize
func (p *Person) DeepCopySerialize() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}
