package kit

import (
	"github.com/go-vgo/robotgo"
)

type Point struct {
	X, Y int
}

// GetColor get the point color
func (p Point) GetColor() string {
	return robotgo.PadHex(robotgo.GetPxColor(p.X, p.Y))
}
