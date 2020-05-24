package kit

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/wilon/robotg"
)

// Help for param:
//
// * color robotg.CHex, is HEX color like this: 0xFF00DD
//
// * imgbitmap robotg.CBitmap, only support .bmp now.
//
//   - robotg.CBitmap(robotg.OpenBitmap("/your/path/image.bmp", 2))
//
//   - or use tool transimg, `go get github.com/vcgo/kit/transimg`:
//      Generate: 	Put .bmp images path to GOPATH, then on the path run `transimg`
//      Use: 		import (. "github.com/your/path/bmpimages")
//           		Screen.FindPic(ImgStr["bmpimages/image.bmp"], 0.14)

// FindColor from area, return nil is success
func (area Area) FindColor(color robotg.CHex, tolerance float64) (Point, error) {
	whereBitmap := robotg.CaptureScreen(area.X, area.Y, area.W, area.H)
	x, y := robotg.FindColor(color, whereBitmap, tolerance)
	robotg.FreeBitmap(whereBitmap)
	if x > 0 || y > 0 {
		return Point{area.X + x, area.Y + y}, nil
	} else {
		return Point{x, y}, errors.New("Cant find color")
	}
}

// FindPic find the position of image from area, return nil is success
func (area Area) FindPic(bmp Bitmap, tolerance float64) (Point, error) {
	findBitmap := robotg.ToCBitmap(robotg.Bitmap(bmp))
	whereBitmap := robotg.CaptureScreen(area.X, area.Y, area.W, area.H)
	x, y := robotg.FindBitmap(findBitmap, whereBitmap, tolerance)
	robotg.FreeBitmap(whereBitmap)
	// robotg.FreeBitmap(findBitmap)
	if x > 0 || y > 0 {
		return Point{area.X + x, area.Y + y}, nil
	} else {
		return Point{-1, -1}, errors.New("Cant find pic")
	}
}

// UntilFindPic do something until find the pic
func (area Area) UntilFindPic(BeforFunc func(i int), bmp Bitmap, tolerance float64) (Point, error) {
	for i := 0; i < 188; i++ {
		p, err := area.FindPic(bmp, tolerance)
		if err == nil {
			return p, nil
		}
		BeforFunc(i)
	}
	return Point{0, 0}, errors.New("Find pic timeout!")
}

// Center get the area center point.
func (area Area) Center() Point {
	return Point{area.X + area.W/2, area.Y + area.H/2}
}

// Start start point
func (area Area) Start() Point {
	return Point{area.X, area.Y}
}

// End end point
func (area Area) End() Point {
	return Point{area.X + area.W, area.Y + area.H}
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

// Spl is a data.
type Spl struct {
	R, C, M, N uint
}

// Split
func (area Area) Split(sp Spl) Area {
	return area.Splice(sp.R, sp.C)[sp.M][sp.N]
}

// 只能往下往右
func (a Area) PointSplit(dist string, p Point) Area {
	switch dist {
	case "right":
		// 即右
		return Area{p.X + 2, a.Y, a.X + a.W - p.X - 2, a.H}
	case "rightDown":
		// 即右下
		return Area{p.X + 2, p.Y + 2, a.X + a.W - p.X - 2, a.Y + a.H - p.Y - 2}
	case "up":
		// 上
		return Area{a.X, p.Y - 2, a.W, a.Y + a.H - p.Y + 2}
	case "down":
		// 即下
		return Area{a.X, p.Y + 2, a.W, a.Y + a.H - p.Y - 2}
	default:
		// 即下
		return Area{a.X, p.Y + 2, a.W, a.Y + a.H - p.Y - 2}
	}
}

// FindPicSeries
// FindFunc true over; false continue
func (area Area) FindPicSeries(dist string, FindFunc func(p Point) bool, bmp Bitmap, tolerance float64) (Point, error) {
	a := area
	for i := 0; i < 188; i++ {
		p, err := a.FindPic(bmp, tolerance)
		if err != nil {
			return Point{-1, -1}, err
		}
		if FindFunc(p) == true {
			return p, err
		} else {
			a = a.PointSplit(dist, p)
			Sleep(233 + rand.Intn(233))
		}
	}
	return Point{0, 0}, errors.New("Find pic timeout!")
}

// Test can save Area to imgage for debug.
func (area Area) Test(pre, path string) {
	path = strings.TrimRight(path, "/") + "/"
	Mkdirs(path)
	pngName := path
	pngName += string(time.Now().Format("2006_01_02.15_04_05")) + "-"
	pngName += pre + "-"
	pngName += "x" + strconv.Itoa(area.X) + "_"
	pngName += "y" + strconv.Itoa(area.Y) + "_"
	pngName += "w" + strconv.Itoa(area.W) + "_"
	pngName += "h" + strconv.Itoa(area.H) + ".png"
	whereBitmap := robotg.CaptureScreen(area.X, area.Y, area.W, area.H)
	_ = robotg.SaveBitmap(whereBitmap, pngName)
	robotg.FreeBitmap(whereBitmap)
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
