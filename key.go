package kit

import (
	"math/rand"

	"github.com/go-vgo/robotgo"
)

var KeyUpMap = make(map[string]bool)

// For the param 'key string',You can refer to:
// https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

// KeyPress is press key func
func KeyPress(key string) {
	KeyDown(key)
	Sleep(44 + rand.Intn(10))
	KeyUp(key)
}

// KeyDown is press down a key
func KeyDown(key string) {
	robotgo.KeyToggle(key, "down")
	KeyUpMap[key] = true
}

// KeyUp is press up a key
func KeyUp(key string) {
	robotgo.KeyToggle(key, "up")
	KeyUpMap[key] = false
}

func KeyDefer() {
	for key, isUp := range KeyUpMap {
		if isUp == true {
			KeyUp(key)
		}
	}
}

// KeyOutput output text
func KeyOutput(text string) {

}
