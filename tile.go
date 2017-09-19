// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package tmx

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

const (
	// Flags representing horizontal, vertical, and diagonal tile flipping.
	// These can be used in combination with gid's, like so:
	//
	// if (gid & flippedHorizontallyFlag) > 0 {
	//     ...draw the tiled flipped horizontally...
	// }
	flippedHorizontallyFlag uint32 = 0x80000000
	flippedVerticallyFlag   uint32 = 0x40000000
	flippedDiagonallyFlag   uint32 = 0x20000000
)

// Tile represents a single tile definition and it's properties
type Tile struct {
	// The ID of the tile
	ID int

	// An array defining the terrain type of each corner of the tile, as indices
	// into the terrain types slice of the tileset this tile came from, in the
	// order of: top left, top right, bottom left, bottom right.
	//
	// -1 values have a meaning of 'no terrain'.
	Terrain [4]int

	// Percentage chance indicating the probability that this tile is chosen
	// when editing with the terrain tool.
	Probability float64

	// Map of properties for the tile
	Properties map[string]string

	// Image for the tile
	Image *Image
}

// String returns a string representation of this tileset.
func (t *Tile) String() string {
	return fmt.Sprintf("Tile(ID=%v)", t.ID)
}
