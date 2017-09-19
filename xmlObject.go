package tmx

type xmlObject struct {
	Name       string        `xml:"name,attr"`
	Type       string        `xml:"type,attr"`
	X          int           `xml:"x,attr"`
	Y          int           `xml:"y,attr"`
	Width      int           `xml:"width,attr"`
	Height     int           `xml:"height,attr"`
	Rotation   float64       `xml:"rotation,attr"`
	Gid        int           `xml:"gid,attr"`
	Visible    int           `xml:"visible,attr"`
	Properties xmlProperties `xml:"properties"`
	Ellipse    *string       `xml:"ellipse"`
	Polygon    xmlPolyset    `xml:"polygon"`
	Polyline   xmlPolyset    `xml:"polyline"`

	// FIXME: alledgedly, object tags can have images under them, but it's not
	// clear what that would mean. There also is no way to create one in
	// Tiled-Qt that I can find anywhere. My guess: it's not supposed to be
	// there at all and none actually have them.
}

func (x xmlObject) toValue() interface{} {
	if x.Ellipse != nil {
		return &Ellipse{
			X:      x.X,
			Y:      x.Y,
			Width:  x.Width,
			Height: x.Height,
		}
	}
	if len(x.Polygon.Data) > 0 {
		return &Polygon{
			X:      x.X,
			Y:      x.Y,
			Points: x.Polygon.toPoints(),
		}
	}
	if len(x.Polyline.Data) > 0 {
		return &Polyline{
			X:      x.X,
			Y:      x.Y,
			Points: x.Polyline.toPoints(),
		}
	}
	return nil
}

func (x xmlObject) toObject() *Object {
	return &Object{
		Name:       x.Name,
		Type:       x.Type,
		X:          x.X,
		Y:          x.Y,
		Width:      x.Width,
		Height:     x.Height,
		Rotation:   x.Rotation,
		Gid:        uint32(x.Gid),
		Visible:    x.Visible != 0,
		Properties: x.Properties.toMap(),
		Value:      x.toValue(),
	}
}
