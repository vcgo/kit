package kit

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/go-ozzo/ozzo-log"
	"github.com/go-vgo/robotgo"
)

// Area is a screen area,
// X,Y is the start point
// W,H is the area's width and hight
type Area struct {
	X, Y, W, H int
}

var (
	Logger      log.Logger
	LogFileName = ""
	Screen      Area
)

func init() {
	w, h := robotgo.GetScreenSize()
	Screen = Area{0, 0, w, h}
}

// InitLogger the func Log() will initialize it.
func InitLogger() {
	if Logger.Category != "" {
		return
	}
	loggerTmp := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	if LogFileName == "" {
		t2.FileName = "app" + string(time.Now().Format(".2006-01-02")) + ".log"
	} else {
		t2.FileName = LogFileName
	}
	t2.MaxLevel = log.LevelNotice
	loggerTmp.Targets = append(loggerTmp.Targets, t1, t2)
	loggerTmp.Open()
	Logger = *loggerTmp
	Logger.Info("InitLogger")
}

// Sleep wait x millisecond
func Sleep(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}

// Fmt is print any variable
func Fmt(desc string, args ...interface{}) string {
	fmt.Println(Logger.Category, desc, args)
	return ""
}

// Log is write log and output any variable easily.
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

// Exit exit the program.
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
	os.Exit(1)
}

// Mkdirs
func Mkdirs(imgpath string) error {
	_, err := os.Stat(imgpath)
	if err == nil {
		return err
	}
	errm := os.MkdirAll(imgpath, 0755)
	if errm == nil {
		return errm
	}
	return nil
}
