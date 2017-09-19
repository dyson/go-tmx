package tmx

type xmlTileset struct {
	Raw          []byte `xml:",innerxml"`
	Firstgid     uint32 `xml:"firstgid,attr"`
	Source       string `xml:"source,attr"`
	Name         string `xml:"name,attr"`
	TileWidth    int    `xml:"tilewidth,attr"`
	TileHeight   int    `xml:"tileheight,attr"`
	Spacing      int    `xml:"spacing,attr"`
	Margin       int    `xml:"margin,attr"`
	Tileoffset   xmlTileoffset
	Properties   xmlProperties   `xml:"properties"`
	Image        xmlImage        `xml:"image"`
	Tile         []xmlTile       `xml:"tile"`
	Terraintypes xmlTerraintypes `xml:"terraintypes"`
}

func (x *xmlTileset) tilesMap() map[int]*Tile {
	tiles := make(map[int]*Tile, len(x.Tile))
	for _, xt := range x.Tile {
		tiles[xt.ID] = xt.toTile()
	}
	return tiles
}

func (x *xmlTileset) terrainTypes() []TerrainType {
	terrainTypes := make([]TerrainType, len(x.Terraintypes.Terrain))
	for i, xt := range x.Terraintypes.Terrain {
		terrainTypes[i] = TerrainType{
			Name: xt.Name,
			Tile: xt.Tile,
		}
	}
	return terrainTypes
}
