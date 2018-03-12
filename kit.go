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

var (
	Logger      log.Logger
	LogFileName = ""
)

func InitLogger() {
	if Logger.Category != "" {
		return
	}
	loggerTmp := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	if LogFileName == "" {
		t2.FileName = "app" + string(time.Now().Format(".2006-01-02")) + ".log"
	}
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
	fmt.Println(Logger.Category, desc, args)
	return ""
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

// Do something until find the pic
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
