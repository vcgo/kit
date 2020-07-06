package kit

import (
	"math/rand"
)

// MoveTo mouse move to x, y
func MoveTo(x int, y int) {
	moveTo(x, y)
}

func MouseToggle(upDown, leftRight string) {
	mouseToggle(upDown, leftRight)
}

// LeftClick mouse left click
func LeftClick() {
	mouseToggle("down", "left")
	Sleep(55 + rand.Intn(10))
	mouseToggle("up", "left")
}

// RightClick mouse right click
func RightClick() {
	mouseToggle("down", "right")
	Sleep(55 + rand.Intn(10))
	mouseToggle("up", "right")
}

// LeftDoubleClick mouse left double click
func LeftDoubleClick() {
	mouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	mouseToggle("up", "left")
	Sleep(99 + rand.Intn(22))
	mouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	mouseToggle("up", "left")
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

func Scroll(dist string) {
	if dist != "up" {
		dist = "down"
	}
	mouseWheel(dist)
	Sleep(88 + rand.Intn(10))
}
