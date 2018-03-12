package kit

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type Area struct {
	X, Y, W, H int
}

var Screen Area

func init() {
	w, h := robotgo.GetScreenSize()
	Screen = Area{0, 0, w, h}
}

func (area Area) Center() (int, int) {
	return area.X + area.W/2, area.Y + area.H/2
}

// Splice Area
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

// test *Area
func (area Area) Test(path string) {

	path = strings.TrimRight(path, "/") + "/"

	Mkdirs(path)
	pngName := path
	pngName += string(time.Now().Format("2006_01_02.15_04_05")) + "-"
	pngName += strconv.Itoa(area.X) + "-" + strconv.Itoa(area.Y) + "-"
	pngName += strconv.Itoa(area.W) + "-" + strconv.Itoa(area.H) + ".png"

	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	_ = robotgo.SaveBitmap(whereBitmap, pngName)
}
