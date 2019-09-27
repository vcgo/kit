package kit

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestBitmap(t *testing.T) {
	bitmap := Screen.Capture()
	fmt.Println(bitmap)
}

func TestFindPic(t *testing.T) {
	area := Screen
	bmp := robotgo.CaptureScreen(100, 100, 20, 20)
	cbmp := robotgo.CBitmap(bmp)
	str := robotgo.TostringBitmap(bmp)
	cbmp = robotgo.CBitmap(robotgo.BitmapStr(str))
	fmt.Println(str)
	x, y, err := area.FindPic(cbmp, 0.1)
	fmt.Println(x, y, err)

	cbmp = robotgo.CBitmap(robotgo.BitmapStr(str))
	area.FindPic(cbmp, 0.1)
	cbmp = robotgo.CBitmap(robotgo.BitmapStr(str))
	area.FindPic(cbmp, 0.1)
	cbmp = robotgo.CBitmap(robotgo.BitmapStr(str))
	area.FindPic(cbmp, 0.1)
	if err != nil {
		t.Error("kit.FindPic coun't find!")
	}
	if x != 100 || y != 100 {
		t.Error("kit.FindPic find results error.")
	}
}
