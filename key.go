package kit

import (
	"math/rand"
)

var KeyUpMap = make(map[string]bool)

// KeyPress is press key func
func KeyPress(key string) {
	KeyDown(key)
	Sleep(55 + rand.Intn(10))
	KeyUp(key)
}

// KeyDown is press down a key
func KeyDown(key string) {
	keyDown(key)
	KeyUpMap[key] = true
}

// KeyUp is press up a key
func KeyUp(key string) {
	keyUp(key)
	KeyUpMap[key] = false
}

// KeyDefer some key is pressing, then use this defer func to up it.
func KeyDefer() {
	for key, needUp := range KeyUpMap {
		if needUp == true {
			KeyUp(key)
			Sleep(55 + rand.Intn(10))
		}
	}
}

// KeyOutput output text
func KeyOutput(text string) {

}
