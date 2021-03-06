package kit

import (
	"errors"
	"io"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	adb "github.com/zach-klippenstein/goadb"
)

var (
	// Only support one device.
	client *adb.Adb
	device *adb.Device
)

// OpenAdbDevice
// Only one device use: OpenAdbDevice()
// Open other device use: OpenAdbDevice("device serial")
func OpenAdbDevice(args ...string) error {
	var err error
	client, err = adb.New()
	if err != nil {
		return err
	}
	client.StartServer()
	_, err = client.ServerVersion()
	if err != nil {
		return err
	}
	devices, err := client.ListDevices()
	if err != nil {
		return err
	}
	if len(devices) < 1 {
		return errors.New("Don't have devices")
	}
	serial := devices[0].Serial
	if len(args) == 1 {
		serial = args[0]
	}
	device = client.Device(adb.DeviceWithSerial(serial))
	// Get size
	output, _ := device.RunCommand("wm size")
	reg, _ := regexp.Compile(`(\d+)x(\d+)`)
	regRes := reg.FindStringSubmatch(output)
	if len(regRes) == 3 {
		w, _ := strconv.Atoi(regRes[1])
		h, _ := strconv.Atoi(regRes[2])
		Screen = Area{0, 0, w, h}
	}
	SwichEngine("adb")
	return nil
}

const (
	EV_SYN = 0
	EV_KEY = 1
	EV_REL = 2
	EV_ABS = 3

	SYN_REPORT    = 0
	SYN_CONFIG    = 1
	SYN_MT_REPORT = 2

	ABS_MT_TOUCH_MAJOR = 48
	ABS_MT_TOUCH_MINOR = 49
	ABS_MT_WIDTH_MAJOR = 50
	ABS_MT_WIDTH_MINOR = 51
	ABS_MT_ORIENTATION = 52
	ABS_MT_POSITION_X  = 53
	ABS_MT_POSITION_Y  = 54
	ABS_MT_TOOL_TYPE   = 55
	ABS_MT_BLOB_ID     = 56
	ABS_MT_TRACKING_ID = 57
	ABS_MT_PRESSURE    = 58

	BTN_TOUCH = 330
)

type event struct {
	Event int
	Code  int
	Value int
}

type adbEngine struct {
}

var adbe adbEngine

// Use:
// Tap: point.LeftClick()
// Swip: point.DargTo(dist)
// SwipMore:
// 		point1.MoveTo()
// 		kit.MouseToggle("down", "left")
// 		point2.MoveTo() && kit.Sleep(2333)
// 		point3.MoveTo() && kit.Sleep(2333)
// 		...
// 		kit.MouseToggle("up", "left")
var (
	adbPoint Point
	taped    = false
	evnid    = 90000
	events   []event
)

func (a adbEngine) keyDown(key string) {}

func (a adbEngine) keyUp(key string) {}

func (a adbEngine) moveTo(x int, y int) {
	if taped == true {
		evnPoint(Point{x, y})
	} else {
		adbPoint = Point{x, y}
	}
}

func (a adbEngine) smoothMoveTo(x int, y int) {}

func (a adbEngine) mouseToggle(upDown, leftRight string) {
	if leftRight == "left" {
		if upDown == "down" {
			evnid++
			events = append(events, event{EV_ABS, ABS_MT_TRACKING_ID, evnid})
			events = append(events, event{EV_KEY, BTN_TOUCH, 1})
			taped = true
			if adbPoint.X > -1 {
				evnPoint(adbPoint)
			}
		} else {
			events = append(events, event{EV_ABS, ABS_MT_TRACKING_ID, 0xffffff})
			events = append(events, event{EV_KEY, BTN_TOUCH, 0})
			evnDone()
			sendevent()
			// 事件结束
			taped = false
		}
	}
}

func (a adbEngine) mouseWheel(dist string) {}

func evnDone() {
	events = append(events, event{EV_SYN, SYN_REPORT, 0})
}
func evnPoint(p Point) {
	events = append(events, event{EV_ABS, ABS_MT_POSITION_X, p.X})
	events = append(events, event{EV_ABS, ABS_MT_POSITION_Y, p.Y})
	events = append(events, event{EV_ABS, ABS_MT_TOUCH_MAJOR, 2 + rand.Intn(3)})
	events = append(events, event{EV_ABS, ABS_MT_TOUCH_MINOR, 2 + rand.Intn(3)})
	evnDone()
	sendevent()
}

func sendevent() {
	arr := []string{}
	for _, e := range events {
		str := "sendevent /dev/input/event1 " + e.str()
		eventStr := strings.Join(append(arr, str), " && ")
		if len(eventStr) >= 255 {
			shell := strings.Join(arr, " && ")
			device.RunCommand(shell)
			arr = []string{}
		}
		arr = append(arr, str)
	}
	shell := strings.Join(arr, " && ")
	device.RunCommand(shell)
	events = []event{}
}

func (e event) str() string {
	str := strconv.Itoa(e.Event) + " "
	str += strconv.Itoa(e.Code) + " "
	str += strconv.Itoa(e.Value)
	return str
}

func (a Area) adbCapture() (Bitmap, error) {
	// Get .png output.
	shell := "screencap -p"
	output, _ := device.RunCommand(shell)
	lines := strings.Split(output, "\n")
	if len(lines) <= 2 {
		return Bitmap{}, errors.New("Capture png error.")
	}
	pngStr := strings.Join(lines[2:], "\n")
	// save to .png
	pngName := string(time.Now().Format("2006_01_02.15_04_05"))
	pngName += strconv.Itoa(rand.Intn(99999)) + "_adbadb.png"
	pngReader := strings.NewReader(pngStr)
	file := path.Join(os.TempDir(), pngName)
	out, _ := os.Create(file)
	defer out.Close()
	io.Copy(out, pngReader)
	// get full bitmap
	fullBit := robotgo.OpenBitmap(file)
	// Need clip ?
	if a != Screen {
		bit := robotgo.GetPortion(fullBit, a.X, a.Y, a.W, a.H)
		robotgo.FreeBitmap(fullBit)
		return NewBitmap(robotgo.ToBitmap(bit), file), nil
	} else {
		return NewBitmap(robotgo.ToBitmap(fullBit), file), nil
	}
}
