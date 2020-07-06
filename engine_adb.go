package kit

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

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
	if len(args) == 0 {
		d := devices[0]
		device = client.Device(adb.DeviceWithSerial(d.Serial))
		SwichEngine("adb")
	} else {
		serial := args[0]
		for _, d := range devices {
			if d.Serial == serial {
				device = client.Device(adb.DeviceWithSerial(d.Serial))
				SwichEngine("adb")
				return nil
			}
		}
		return errors.New("Don't have serial: " + serial)
	}
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
	taped    bool
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
