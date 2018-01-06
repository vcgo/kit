package kit

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
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

var (
	Logger log.Logger
)

func init() {
	loggerTmp := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = "app.log"
	t2.MaxLevel = log.LevelNotice
	loggerTmp.Targets = append(loggerTmp.Targets, t1, t2)
	loggerTmp.Open()
	Logger = *loggerTmp
	defer Logger.Close()
}

// sleep Millisecond
func Sleep(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}

// write log
func Log(desc string, args ...interface{}) {
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := strings.Split(file, "src/")[1]
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	Logger.Notice(fileLine + "  " + desc + argsDesc)
}

// Exit
func Exit(args ...interface{}) {
	argsDesc := ""
	for _, val := range args {
		argsDesc += " " + strings.TrimSpace(fmt.Sprintln(val))
	}
	_, file, line, _ := runtime.Caller(1)
	fileName := strings.Split(file, "src/")[1]
	fileLine := fileName + " " + strconv.Itoa(line) + " "
	if argsDesc != "" {
		Logger.Error(fileLine + "  " + argsDesc)
	}

	Logger.Error(fileLine + " Exit!!!\n")
	os.Exit(1)
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
func FindColor(area Area, color robotgo.CHex) (int, int, error) {
	x, y := robotgo.FindColorCS(area.X, area.Y, area.W, area.H, color, 0)

	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return x, y, errors.New("Cant find color")
	}
}

// find pic from area, return nil is success
func FindPic(area Area, imgbitmap robotgo.CBitmap, tolerance float32) (int, int, error) {

	whereBitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	findBitmap := robotgo.ToMMBitmapRef(imgbitmap)
	x, y := robotgo.FindBitmap(findBitmap, whereBitmap, tolerance)

	if x > 0 || y > 0 {
		return area.X + x, area.Y + y, nil
	} else {
		return -1, -1, errors.New("Cant find pic")
	}
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
