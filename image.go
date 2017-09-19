package tmx

import (
	"fmt"
	"image/color"
)

// Image represents the source and properties of a image
type Image struct {
	Format string
	Source string
	Trans  color.RGBA
	Width  int
	Height int
}

// String returns a string representation of this image
func (i *Image) String() string {
	return fmt.Sprintf("Image(Format=%q, Source=%q, Size=%dx%dpx)", i.Format, i.Source, i.Width, i.Height)
}
