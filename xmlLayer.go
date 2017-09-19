package tmx

type xmlLayer struct {
	Name       string        `xml:"name,attr"`
	X          int           `xml:"x,attr"`
	Y          int           `xml:"y,attrr"`
	Width      int           `xml:"width,attrr"`
	Height     int           `xml:"height,attrr"`
	Opacity    float64       `xml:"opacity,attr"`
	Visible    int           `xml:"visible,attr"`
	OffsetX    int           `xml:"offsetx,attr"`
	OffsetY    int           `xml:"offsety,attr"`
	Properties xmlProperties `xml:"properties"`
	Data       xmlData       `xml:"data"`
}

// TODO: width and height?
func (x xmlLayer) toLayer(width, height int) (*Layer, error) {
	tiles, err := x.Data.tiles(width, height)
	if err != nil {
		return nil, err
	}
	return &Layer{
		Name:       x.Name,
		X:          x.X,
		Y:          x.Y,
		Width:      x.Width,
		Height:     x.Height,
		Opacity:    x.Opacity,
		Visible:    x.Visible != 0,
		OffsetX:    x.OffsetX,
		OffsetY:    x.OffsetY,
		Properties: x.Properties.toMap(),
		Tiles:      tiles,
	}, nil
}
