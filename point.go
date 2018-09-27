package kit

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

type Point struct {
	X, Y int
}

// GetColor get the point color
func (p Point) GetColor() string {
	return robotgo.PadHex(robotgo.GetPxColor(p.X, p.Y))
}

// DragTo start on a point, drag to another point.
func (p Point) DragTo(dstp Point) {
	MoveTo(p.X, p.Y)
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("down", "left")
	Sleep(55 + rand.Intn(10))
	MoveTo(dstp.X, dstp.Y)
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("up", "left")
	Sleep(55 + rand.Intn(10))
}
