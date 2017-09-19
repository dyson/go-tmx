// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package tmx

import (
	"fmt"
	"image/color"
)

// ObjectGroup represents a group of objects.
type ObjectGroup struct {
	// The name of this object group.
	Name string

	// Color of this object group.
	Color color.RGBA

	// Value between 0 and 1 representing the opacity of the object group.
	Opacity float64

	// Boolean value representing whether or not the object group is visible.
	Visible bool

	// Map of properties for this object group.
	Properties map[string]string

	// List of objects in this object group.
	Objects []*Object
}

// String returns a string representation of this object group, like:
//  ObjectGroup(Name="the name", 500 objects)
func (o *ObjectGroup) String() string {
	return fmt.Sprintf("ObjectGroup(Name=%q, %d objects)", o.Name, len(o.Objects))
}
