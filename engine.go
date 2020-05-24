package kit

import "github.com/go-vgo/robotgo"

func keyDown(key string) {
	robotgo.KeyToggle(key, "down")
}

func keyUp(key string) {
	robotgo.KeyToggle(key, "up")
}

func moveTo(x int, y int) {
	robotgo.Move(x, y)
}

func smoothMoveTo(x int, y int) {
	robotgo.MoveMouseSmooth(x, y)
}

func mouseToggle(upDown string, leftRight string) {
	robotgo.MouseToggle(upDown, leftRight)
}

func ScrollMouse(dist string) {
	robotgo.ScrollMouse(1, dist)
}
