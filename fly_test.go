package kit

import (
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	err := OpenFly()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer CloseFly()
	// moveTo
	MoveTo(100, 300)
	Sleep(666)
	MouseToggle("left", "down")
	Sleep(666)
	MoveTo(500, 500)
	Sleep(666)
	MouseToggle("left", "up")
}
