package kit

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

// For the param 'key string',You can refer to:
// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

// KeyPress is press key func
func KeyPress(key string) {
	robotgo.KeyToggle(key, "down")
	Sleep(55 + rand.Intn(10))
	robotgo.KeyToggle(key, "up")
}

// KeyDown is press down a key
func KeyDown(key string) {
	robotgo.KeyToggle(key, "down")
}

// KeyUp is press up a key
func KeyUp(key string) {
	robotgo.KeyToggle(key, "up")
}
