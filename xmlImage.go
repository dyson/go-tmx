package tmx

type xmlImage struct {
	Format string `xml:"format,attr"`
	// TODO: expose image <data> element for embedded image data ?
	Source string `xml:"source,attr"`
	Trans  string `xml:"trans,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
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
