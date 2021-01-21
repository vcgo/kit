package kit

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

// Area is a screen area,
// X,Y is the start point
// W,H is the area's width and hight
type Area struct {
	X, Y, W, H int
}

func A(x, y, w, h int) Area {
	return Area{x, y, w, h}
}

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
func (area Area) FindColor(color robotgo.CHex, tolerance float64) (Point, error) {
	var x, y int
	if engine == "adb" {
		bitmap, _ := area.adbCapture()
		whereBitmap := robotgo.ToCBitmap(bitmap.Bitmap)
		x, y = robotgo.FindColor(color, whereBitmap, tolerance)
		bitmap.Free()
	} else {
		whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
		x, y = robotgo.FindColor(color, whereBitmap, tolerance)
		robotgo.FreeBitmap(whereBitmap)
	}
	if x > 0 || y > 0 {
		return Point{area.X + x, area.Y + y}, nil
	} else {
		return Point{x, y}, errors.New("Cant find color")
	}
}

// FindPic find the position of image from area, return nil is success.
// bmp unusefull need to free.
func (area Area) FindPic(bmp Bitmap, tolerance float64) (Point, error) {
	var x, y int
	findBitmap := robotgo.ToCBitmap(bmp.Bitmap)
	if engine == "adb" {
		bitmap, _ := area.adbCapture()
		whereBitmap := robotgo.ToCBitmap(bitmap.Bitmap)
		x, y = robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)
		bitmap.Free()
	} else {
		whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
		x, y = robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)
		robotgo.FreeBitmap(whereBitmap)
	}
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
	return Point{area.X + int(area.W/2), area.Y + int(area.H/2)}
}

// Start start point
func (area Area) Start() Point {
	return Point{area.X, area.Y}
}

// End end point
func (area Area) End() Point {
	return Point{area.X + area.W - 1, area.Y + area.H - 1}
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

// PointSplit
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
func (area Area) FindPicSeries(
	dist string,
	FindFunc func(p Point) bool,
	bmp Bitmap,
	tolerance float64,
) (Point, error) {
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
	bm := area.Capture()
	bm.SavePng(pngName)
	bm.Free()
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

// HexPoint
// HexMatrix find this Point Hex
type HexPoint struct {
	Point Point
	Hex   int32
}

// HexArea
// HexMatrix find this Area count Hex >= Count
type HexArea struct {
	Area  Area
	Hex   int32
	Count int
}

// HexMatrix
// A point or color matrix for find.
type HexMatrix struct {
	Hex    uint32
	Points []HexPoint
	Areas  []HexArea
}

// FindHexMatrix
func (area Area) FindHexMatrix(hm HexMatrix) (Point, error) {
	// param
	tol := 0.0
	chex := robotgo.UintToHex(hm.Hex)
	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	defer robotgo.FreeBitmap(whereBitmap)
	// is point match
	var isMatch = func(p Point) bool {
		match := true
		// find points
		for _, hp := range hm.Points {
			m, n := p.X+hp.Point.X-area.X, p.Y+hp.Point.Y-area.Y
			if m > area.W {
				match = false
				break
			}
			if n > area.H {
				match = false
				break
			}
			hex := robotgo.CHex(robotgo.GetColor(whereBitmap, m, n))
			if hp.Hex >= 0 {
				if hex != robotgo.UintToHex(uint32(hp.Hex)) {
					match = false
					break
				}
			} else {
				if hex == robotgo.UintToHex(uint32(-hp.Hex)) {
					match = false
					break
				}
			}
		}
		// count areas
		for _, ha := range hm.Areas {
			a := ha.Area
			hmBmp := robotgo.GetPortion(whereBitmap, p.X+a.X-area.X, p.Y+a.Y-area.Y, a.W, a.H)
			defer robotgo.FreeBitmap(hmBmp)
			count := 0
			if ha.Hex >= 0 {
				color := robotgo.UintToHex(uint32(ha.Hex))
				count = robotgo.CountColor(color, hmBmp, tol)
			} else {
				color := robotgo.UintToHex(uint32(-ha.Hex))
				nc := robotgo.CountColor(color, hmBmp, tol)
				count = a.W*a.H - nc
			}
			if count < ha.Count {
				match = false
				break
			}
		}
		// res
		return match
	}
	// search
	restArea := area
	for {
		rx, ry, rw, rh := restArea.X-area.X, restArea.Y-area.Y, restArea.W, restArea.H
		if rx < 0 || ry < 0 || rw <= 0 || rh <= 0 || rx+rw > area.W || ry+rh > area.H {
			break
		}
		restBitmap := robotgo.GetPortion(whereBitmap, rx, ry, rw, rh)
		defer robotgo.FreeBitmap(restBitmap)
		x, y := robotgo.FindColor(chex, restBitmap, tol)
		if x < 0 {
			break
		}
		// line
		lineArea := A(restArea.X+x, restArea.Y+y, restArea.W-x, 1)
		for l := lineArea.X; l < lineArea.X+lineArea.W; l++ {
			p := P(l, lineArea.Y)
			if isMatch(p) {
				return p, nil
			}
		}
		// area
		restArea = A(restArea.X, restArea.Y+y+1, restArea.W, restArea.H-y-1)
	}
	return Point{-1, -1}, errors.New("Cant find HexMatrix")
}

// FindHexMatrixGo
// Gorutine, Slowly, high CPU fun.
func (area Area) FindHexMatrixGo(hm HexMatrix) (Point, error) {
	// param
	tol := 0.0
	chex := robotgo.UintToHex(hm.Hex)
	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	defer robotgo.FreeBitmap(whereBitmap)
	wg := &sync.WaitGroup{}
	resCh := make(chan Point)
	limitCh := make(chan int, 100)
	// is point match
	var isMatch = func(p Point) bool {
		match := true
		// find points
		for _, hp := range hm.Points {
			m, n := p.X+hp.Point.X-area.X, p.Y+hp.Point.Y-area.Y
			if m > area.W {
				match = false
				break
			}
			if n > area.H {
				match = false
				break
			}
			hex := robotgo.CHex(robotgo.GetColor(whereBitmap, m, n))
			if hp.Hex >= 0 {
				if hex != robotgo.UintToHex(uint32(hp.Hex)) {
					match = false
					break
				}
			} else {
				if hex == robotgo.UintToHex(uint32(-hp.Hex)) {
					match = false
					break
				}
			}
		}
		// count areas
		for _, ha := range hm.Areas {
			a := ha.Area
			hmBmp := robotgo.GetPortion(whereBitmap, p.X+a.X, p.Y+a.Y, a.W, a.H)
			defer robotgo.FreeBitmap(hmBmp)
			count := 0
			if ha.Hex >= 0 {
				color := robotgo.UintToHex(uint32(ha.Hex))
				count = robotgo.CountColor(color, hmBmp, tol)
			} else {
				color := robotgo.UintToHex(uint32(-ha.Hex))
				nc := robotgo.CountColor(color, hmBmp, tol)
				count = a.W*a.H - nc
			}
			if count < ha.Count {
				match = false
				break
			}
		}
		// res
		return match
	}
	// search
	restArea := area
	for {
		rx, ry, rw, rh := restArea.X-area.X, restArea.Y-area.Y, restArea.W, restArea.H
		if rx < 0 || ry < 0 || rw <= 0 || rh <= 0 || rx+rw > area.W || ry+rh > area.H {
			break
		}
		restBitmap := robotgo.GetPortion(whereBitmap, rx, ry, rw, rh)
		defer robotgo.FreeBitmap(restBitmap)
		x, y := robotgo.FindColor(chex, restBitmap, tol)
		if x < 0 {
			break
		}
		// line
		lineArea := A(restArea.X+x, restArea.Y+y, restArea.W-x, 1)
		for l := lineArea.X; l < lineArea.X+lineArea.W; l++ {
			p := P(l, lineArea.Y)
			wg.Add(1)
			limitCh <- 1
			go func(p Point, resCh chan Point) {
				defer func() {
					wg.Done()
					<-limitCh
				}()
				if isMatch(p) {
					resCh <- p
				}
			}(p, resCh)
		}
		// area
		restArea = A(restArea.X, restArea.Y+y+1, restArea.W, restArea.H-y-1)
	}
	go func() {
		wg.Wait()
		resCh <- Point{-1, -1}
	}()
	for {
		select {
		case p := <-resCh:
			if p.X < 0 {
				return Point{-1, -1}, errors.New("Cant find HexMatrix")
			}
			return p, nil
		}
	}
}
