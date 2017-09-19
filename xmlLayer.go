package tmx

type xmlLayer struct {
	Name    string  `xml:"name,attr"`
	Opacity float64 `xml:"opacity,attr"`
	Visible int     `xml:"visible,attr"`
	Data    xmlData `xml:"data"`
}

func (x xmlLayer) toLayer(width, height int) (*Layer, error) {
	tiles, err := x.Data.tiles(width, height)
	if err != nil {
		return nil, err
	}
	return &Layer{
		Name:    x.Name,
		Opacity: x.Opacity,
		Visible: x.Visible != 0,
		Tiles:   tiles,
	}, nil
}
