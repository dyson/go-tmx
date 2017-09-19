package tmx

type xmlImage struct {
	Format string `xml:"format,attr"`
	//ID deprecated
	Source string `xml:"source,attr"`
	Trans  string `xml:"trans,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	//Data not possible to create using Tiled
}

func (x xmlImage) toImage() *Image {
	return &Image{
		Format: x.Format,
		Source: x.Source,
		Trans:  hexToRGBA(x.Trans),
		Width:  x.Width,
		Height: x.Height,
	}
}
