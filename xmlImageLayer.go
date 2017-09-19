package tmx

type xmlImageLayer struct {
	Name    string `xml:"name,attr"`
	OffsetX int    `xml:"offsetx,attr"`
	OffsetY int    `xml:"offsety,attr"`
	//X deprecated
	//Y deprecated
	Opacity    float64       `xml:"opacity,attr"`
	Visible    int           `xml:"visible,attr"`
	Properties xmlProperties `xml:"properties"`
	Image      xmlImage      `xml:"image"`
}

func (x xmlImageLayer) toImageLayer() *ImageLayer {
	return &ImageLayer{
		Name:       x.Name,
		OffsetX:    x.OffsetX,
		OffsetY:    x.OffsetY,
		Opacity:    x.Opacity,
		Visible:    x.Visible != 0,
		Properties: x.Properties.toMap(),
		Image:      x.Image.toImage(),
	}
}
