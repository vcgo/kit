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
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("down", "left")
	Sleep(66 + rand.Intn(22))
	MoveTo(dstp.X, dstp.Y)
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("up", "left")
	Sleep(66 + rand.Intn(22))
}

// RightClick mouse right click
func (p Point) RightClick() {
	robotgo.Move(p.X, p.Y)
	Sleep(33 + rand.Intn(33))
	robotgo.MouseToggle("down", "right")
	Sleep(33 + rand.Intn(33))
	robotgo.MouseToggle("up", "right")
}
