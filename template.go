package tmx

import (
	"fmt"
)

// Template represents a saved map object
type Template struct {
	Name   string
	ID     int
	Object *Object
}

// String returns a string representation of this template
func (t *Template) String() string {
	return fmt.Sprintf("Template(Name=%q, ID=%d)", t.Name, t.ID)
}
