package kit

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

// MoveTo mouse move to x, y
func MoveTo(x int, y int) {
	robotgo.Move(x, y)
	// robotgo.MoveSmooth(x, y, 0.033, 0.046, rand.Intn(9))
}

// MoveSmoothTo mouse move to x, y
func MoveSmoothTo(x int, y int) {
	robotgo.MoveSmooth(x, y, 0.033, 0.046, rand.Intn(9))
}

// LeftClick mouse left click
func LeftClick() {
	robotgo.MouseToggle("down", "left")
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("up", "left")
}

// RightClick mouse right click
func RightClick() {
	robotgo.MouseToggle("down", "right")
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("up", "right")
}

// LeftDoubleClick mouse left double click
func LeftDoubleClick() {
	robotgo.MouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("up", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("up", "left")
	Sleep(99 + rand.Intn(22))
}

// MoveClick mouse move to and left click
func MoveClick(x int, y int) {
	MoveTo(x, y)
	Sleep(55 + rand.Intn(22))
	LeftClick()
}

// MoveDoubleClick mouse move to and double click
func MoveDoubleClick(x int, y int) {
	MoveTo(x, y)
	Sleep(88 + rand.Intn(10))
	LeftDoubleClick()
}
