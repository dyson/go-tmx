package tmx

// NOTE: x, y, width and height attributes are apparently meaningless:
//  https://github.com/bjorn/tiled/wiki/TMX-Map-Format#objectgroup
type xmlObjectgroup struct {
	Name       string        `xml:"name,attr"`
	Color      string        `xml:"color,attr"`
	Opacity    float64       `xml:"opacity,attr"`
	Visible    int           `xml:"visible,attr"`
	Properties xmlProperties `xml:"properties"`
	Object     []xmlObject   `xml:"object"`
}

func (x xmlObjectgroup) toObjectGroup() *ObjectGroup {
	objects := make([]*Object, len(x.Object))
	for i, o := range x.Object {
		objects[i] = o.toObject()
	}
	return &ObjectGroup{
		Name:       x.Name,
		Color:      hexToRGBA(x.Color),
		Opacity:    x.Opacity,
		Visible:    x.Visible != 0,
		Properties: x.Properties.toMap(),
		Objects:    objects,
	}
}
