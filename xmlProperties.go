package tmx

type xmlProperties struct {
	Property []xmlProperty `xml:"property"`
}

func (p xmlProperties) toMap() map[string]string {
	m := make(map[string]string, len(p.Property))
	for _, p := range p.Property {
		m[p.Name] = p.Value
	}
	return m
}
