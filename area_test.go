package kit

import (
	"fmt"
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestPointSplit(t *testing.T) {
	area := Area{100, 100, 300, 200}
	area.Test("a", "./tests")
	p := Point{130, 166}
	// right
	a1 := area.PointSplit("right", p)
	a1.Test("a1", "./tests")
	// rightDown
	a2 := area.PointSplit("rightDown", p)
	a2.Test("a2", "./tests")
	// down
	a3 := area.PointSplit("down", p)
	a3.Test("a3", "./tests")
}

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
