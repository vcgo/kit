package kit

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// Help for param:
//
// * color robotgo.CHex, is HEX color like this: 0xFF00DD
//
// * imgbitmap robotgo.CBitmap, only support .bmp now.
//
//   - robotgo.CBitmap(robotgo.OpenBitmap("/your/path/image.bmp", 2))
//
//   - or use tool transimg, `go get github.com/vcgo/kit/transimg`:
//      Generate: 	Put .bmp images path to GOPATH, then on the path run `transimg`
//      Use: 		import (. "github.com/your/path/bmpimages")
//           		Screen.FindPic(ImgStr["bmpimages/image.bmp"], 0.14)

// FindColor from area, return nil is success
func (area Area) FindColor(color robotgo.CHex, tolerance float64) (int, int, error) {
	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	x, y := robotgo.FindColor(color, whereBitmap, tolerance)
	robotgo.FreeBitmap(whereBitmap)
	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return x, y, errors.New("Cant find color")
	}
}

// FindPic find the position of image from area, return nil is success
func (area Area) FindPic(imgbitmap robotgo.CBitmap, tolerance float64) (int, int, error) {
	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	findBitmap := robotgo.ToMMBitmapRef(imgbitmap)
	x, y := robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)
	robotgo.FreeBitmap(whereBitmap)
	// robotgo.FreeBitmap(findBitmap)
	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return -1, -1, errors.New("Cant find pic")
	}
}

// UntilFindPic do something until find the pic
func (area Area) UntilFindPic(BeforFunc func(), imgbitmap robotgo.CBitmap, tolerance float64) (int, int) {
	for i := 0; i < 188; i++ {
		x, y, err := area.FindPic(imgbitmap, tolerance)
		if err == nil {
			return x, y
		}
		BeforFunc()
	}
	return 0, 0
}

// Center get the area center point.
func (area Area) Center() (int, int) {
	return area.X + area.W/2, area.Y + area.H/2
}

// Splice Area to a arrays.
//
// Such as : 2 row 3 col
// |-----------------|
// | 0,0 | 0,1 | 0,2 |
// |-----|-----|-----|
// | 1,0 | 1,1 | 1,2 |
// |-----------------|
func (area Area) Splice(srow uint, scol uint) [][]Area {
	row := int(srow)
	col := int(scol)
	w := int(math.Ceil(float64(area.W) / float64(col)))
	h := int(math.Ceil(float64(area.H) / float64(row)))
	res := make([][]Area, row)
	for r := 0; r < row; r++ {
		areaArr := make([]Area, col)
		for c := 0; c < col; c++ {
			areaArr[c] = Area{area.X + c*w, area.Y + r*h, w, h}
		}
		res[r] = areaArr
	}
	return res
}

// Test can save Area to imgage for debug.
func (area Area) Test(path string) {
	path = strings.TrimRight(path, "/") + "/"
	Mkdirs(path)
	pngName := path
	pngName += string(time.Now().Format("2006_01_02.15_04_05")) + "-"
	pngName += strconv.Itoa(area.X) + "-" + strconv.Itoa(area.Y) + "-"
	pngName += strconv.Itoa(area.W) + "-" + strconv.Itoa(area.H) + ".png"
	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	_ = robotgo.SaveBitmap(whereBitmap, pngName)
	robotgo.FreeBitmap(whereBitmap)
}

// CountPixel list Area pixel.
func (area Area) CountPixel() map[Point]string {
	res := make(map[Point]string)
	for x := area.X; x < area.X+area.W; x++ {
		for y := area.Y; y < area.Y+area.H; y++ {
			point := Point{x, y}
			res[point] = point.GetColor()
		}
	}
	return res
}
