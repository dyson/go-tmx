package tmx

type xmlMap struct {
	Version         string           `xml:"version,attr"`
	Orientation     string           `xml:"orientation,attr"`
	Width           int              `xml:"width,attr"`
	Height          int              `xml:"height,attr"`
	TileWidth       int              `xml:"tilewidth,attr"`
	TileHeight      int              `xml:"tileheight,attr"`
	BackgroundColor string           `xml:"backgroundcolor,attr"`
	Properties      xmlProperties    `xml:"properties"`
	Tileset         []xmlTileset     `xml:"tileset"`
	Layer           []xmlLayer       `xml:"layer"`
	Objectgroup     []xmlObjectgroup `xml:"objectgroup"`
}
