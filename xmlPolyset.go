package tmx

import (
	"strconv"
	"strings"
)

// Polygon and Polyset are identical, we share definitions here.
type xmlPolyset struct {
	Data string `xml:"points,attr"`
}

func (x xmlPolyset) toPoints() (points []Point) {
	pairs := strings.Split(x.Data, " ")
	points = make([]Point, 0, len(pairs))
	for _, pair := range pairs {
		xy := strings.Split(strings.TrimSpace(pair), ",")
		if len(xy) == 2 {
			x, _ := strconv.Atoi(strings.TrimSpace(xy[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(xy[1]))
			points = append(points, Point{
				X: x,
				Y: y,
			})
		}
	}
	return points
}
