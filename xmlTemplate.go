package tmx

type xmlTemplate struct {
	Name   string    `xml:"name,attr"`
	ID     int       `xml:"id,attr"`
	Object xmlObject `xml:"object"`
}

func (x xmlTemplate) toTemplate() *Image {
	return &Template{
		Name:   x.Name,
		ID:     x.ID,
		Object: x.Object.toObject,
	}
}
