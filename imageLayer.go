package tmx

import (
	"fmt"
)

// ImageLayer represents the source and properties of an image layer
type ImageLayer struct {
	Name       string
	OffsetX    int
	OffsetY    int
	Opacity    float64
	Visible    bool
	Properties map[string]string
	Image      *Image
}

// String returns a string representation of this image layer
func (i *ImageLayer) String() string {
	return fmt.Sprintf(
		"ImageLayer(Name=%q, OffsetX=%d, OffsetY=%d, Opacity=%1.f, Visible=%v)",
		i.Name, i.OffsetX, i.OffsetY, i.Opacity, i.Visible
	)
}
