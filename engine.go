package kit

import "github.com/go-vgo/robotgo"

// engine
// support for fly, adb, default
// default, win/macos support
// fly, a usb hardware support
// adb, android support
var engine = "default"

// SwichEngine
func SwichEngine(kind string) {
	switch kind {
	// Must use kit.OpenFly() success
	// Need fly.dll
	case "fly":
		engine = "fly"
	// Must use kit.OpenAdb() success
	// Need adb shell
	case "adb":
		engine = "adb"
	default:
		engine = "default"
	}
}

func keyDown(key string) {
	switch engine {
	case "fly":
		flye.keyDown(key)
	case "adb":
	default:
		robotgo.KeyToggle(key, "down")
	}
}

func keyUp(key string) {
	switch engine {
	case "fly":
		flye.keyUp(key)
	case "adb":
	default:
		robotgo.KeyToggle(key, "up")
	}
}

func moveTo(x int, y int) {
	switch engine {
	case "fly":
		flye.moveTo(x, y)
	case "adb":
		adbe.moveTo(x, y)
	default:
		robotgo.Move(x, y)
	}
}

func smoothMoveTo(x int, y int) {
	switch engine {
	case "fly":
		flye.smoothMoveTo(x, y)
	case "adb":
		adbe.moveTo(x, y)
	default:
		robotgo.MoveMouseSmooth(x, y)
	}
}

func mouseToggle(upDown, leftRight string) {
	switch engine {
	case "fly":
		flye.mouseToggle(upDown, leftRight)
	case "adb":
		adbe.mouseToggle(upDown, leftRight)
	default:
		robotgo.MouseToggle(upDown, leftRight)
	}
}

func mouseWheel(dist string) {
	switch engine {
	case "fly":
		flye.mouseWheel(dist)
	case "adb":
	default:
		robotgo.ScrollMouse(1, dist)
	}
}
