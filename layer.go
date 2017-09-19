package tmx

import (
	"fmt"
)

// Layer represents a single map layer and all of it's tiles
type Layer struct {
	Name       string
	Opacity    float64
	Visible    bool
	OffsetX    int
	OffsetY    int
	Properties map[string]string
	// A map of 2D coordinates in this layer to so called "global tile IDs"
	// (gids).
	//
	// 2D coordinates whose gid's are zero (I.e. 'no tile') are not stored in
	// the map for efficiency reasons (as a good majority are zero).
	//
	// gids are global, since they may refere to a tile from any of the
	// tilesets used by the map. In order to find out from which tileset the
	// tile is you need to find the tileset with the highest Firstgid that is
	// still lower or equal than the gid. The tilesets are always stored with
	// increasing firstgids.
	Tiles map[Coord]uint32 // from data
}

// String returns a string representation of this layer
func (l *Layer) String() string {
	return fmt.Sprintf("Layer(Name=%q, Opacity=%1.f, Visible=%v)", l.Name, l.Opacity, l.Visible)
}

// Coord represents a single 2D coordinate pair (x, y)
type Coord struct {
	X, Y int
}
