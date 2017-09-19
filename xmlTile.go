package tmx

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

type xmlTile struct {
	ID          int           `xml:"id,attr"`
	Terrain     []byte        `xml:"terrain,attr"`
	Probability float64       `xml:"probability,attr"`
	Properties  xmlProperties `xml:"properties"`
	Image       xmlImage      `xml:"image"`
}

func (x xmlTile) terrainArray() (indices [4]int) {
	var (
		buf = bytes.NewBuffer(x.Terrain)
		k   = 0
	)
	indices = [4]int{-1, -1, -1, -1}
	terrain, err := csv.NewReader(buf).ReadAll()
	if err != nil {
		return
	}
	for _, i := range terrain {
		for _, j := range i {
			index, _ := strconv.Atoi(j)
			indices[k] = index
			k++
			if k >= 4 {
				return
			}
		}
	}
	return
}

func (x xmlTile) toTile() *Tile {
	return &Tile{
		ID:          x.ID,
		Terrain:     x.terrainArray(),
		Probability: x.Probability,
		Properties:  x.Properties.toMap(),
		Image:       x.Image.toImage(),
	}
}
