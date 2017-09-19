package tmx

type xmlLayer struct {
	Name string `xml:"name,attr"`
	//X defaults to 0 and cannot be changed in Tiled
	//Y defaults to 0 and cannot be changed in Tiled
	//Width required but always the same as the map width
	//Height required but always the same as the map height
	Opacity    float64 `xml:"opacity,attr"`
	Visible    int     `xml:"visible,attr"`
	OffsetX    int
	OffsetY    int
	Properties xmlProperties `xml:"properties"`
	Data       xmlData       `xml:"data"`
}

func (x xmlLayer) toLayer(width, height int) (*Layer, error) {
	tiles, err := x.Data.tiles(width, height)
	if err != nil {
		return nil, err
	}
	return &Layer{
		Name:       x.Name,
		Opacity:    x.Opacity,
		Visible:    x.Visible != 0,
		OffsetX:    x.OffsetX,
		OffsetY:    x.OffsetY,
		Properties: x.Properties,
		Tiles:      tiles,
	}, nil
}
