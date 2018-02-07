package kit

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/go-ozzo/ozzo-log"
	"github.com/go-vgo/robotgo"
)

type Area struct {
	X, Y, W, H int
}

type Point struct {
	X, Y int
}

var (
	Logger log.Logger
)

func InitLogger() {
	if Logger.Category != "" {
		return
	}
	loggerTmp := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = "app.log"
	t2.MaxLevel = log.LevelNotice
	loggerTmp.Targets = append(loggerTmp.Targets, t1, t2)
	loggerTmp.Open()
	Logger = *loggerTmp
	Logger.Info("InitLogger")
}

// sleep Millisecond
func Sleep(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}

// write log
func Log(desc string, args ...interface{}) string {
	InitLogger()
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	Logger.Notice(fileLine + "  " + desc + argsDesc)
	return desc + argsDesc
}

// Exit
func Exit(desc string, args ...interface{}) {
	InitLogger()
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := path.Base(file)
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	Logger.Error(fileLine + "  " + desc + argsDesc)
	Logger.Close()
	robotgo.ShowAlert("DNF GO Error!", desc)
	// os.Exit(1)
}

// mouse move to x, y
func MoveTo(x int, y int) {
	robotgo.Move(x, y)
	// robotgo.MoveSmooth(x, y, 900, 1000)
}

// mouse left click
func LeftClick() {
	robotgo.MouseToggle("down", "left")
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("up", "left")
}

// find color from area, return nil is success
func (area Area) FindColor(color robotgo.CHex) (int, int, error) {
	x, y := robotgo.FindColorCS(area.X, area.Y, area.W, area.H, color, 0)

	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return x, y, errors.New("Cant find color")
	}
}

// find pic from area, return nil is success
func (area Area) FindPic(imgbitmap robotgo.CBitmap, tolerance float64) (int, int, error) {

	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	findBitmap := robotgo.ToMMBitmapRef(imgbitmap)
	x, y := robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)

	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return -1, -1, errors.New("Cant find pic")
	}
}

// find pic from area, return nil is success
func GetColor(x int, y int) string {
	// colorStr :=
	// fmt.Println("...", colorStr)
	return robotgo.PadHex(robotgo.GetPxColor(x, y))
}

// test *Area
func (area Area) Test(path string) {

	path = strings.TrimRight(path, "/") + "/"

	Mkdirs(path)
	pngName := path
	pngName += strconv.Itoa(area.X) + "-" + strconv.Itoa(area.Y) + "-"
	pngName += strconv.Itoa(area.W) + "-" + strconv.Itoa(area.H)
	pngName += string(time.Now().Format(".2006-01-02.15_04_05")) + ".png"

	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	_ = robotgo.SaveBitmap(whereBitmap, pngName)
}

// Key Press
func KeyPress(key string) {
	robotgo.KeyToggle(key, "down")
	Sleep(55 + rand.Intn(10))
	robotgo.KeyToggle(key, "up")
}

// Key Down
func KeyDown(key string) {
	robotgo.KeyToggle(key, "down")
}

// Key Up
func KeyUp(key string) {
	robotgo.KeyToggle(key, "up")
}

func Mkdirs(imgpath string) bool {

	_, err := os.Stat(imgpath)
	if err == nil {
		return true
	}
	errm := os.MkdirAll(imgpath, 0755)
	if errm == nil {
		return true
	}
	return false

}
