package kit

import (
	"errors"
	"math/rand"
	"regexp"
	"strconv"
	"syscall"

	"github.com/StackExchange/wmi"
	"github.com/vcgo/win"
)

// FlyKeyCode flye key's code map
var flyKeyCode = map[string]int{
	"a":            4,
	"b":            5,
	"c":            6,
	"d":            7,
	"e":            8,
	"f":            9,
	"g":            10,
	"h":            11,
	"i":            12,
	"j":            13,
	"k":            14,
	"l":            15,
	"m":            16,
	"n":            17,
	"o":            18,
	"p":            19,
	"q":            20,
	"r":            21,
	"s":            22,
	"t":            23,
	"u":            24,
	"v":            25,
	"w":            26,
	"x":            27,
	"y":            28,
	"z":            29,
	"1":            30,
	"2":            31,
	"3":            32,
	"4":            33,
	"5":            34,
	"6":            35,
	"7":            36,
	"8":            37,
	"9":            38,
	"0":            39,
	"enter":        40,
	"escape":       41,
	"backspace":    42,
	"tab":          43,
	"space":        44,
	"-":            45,
	"=":            46,
	"[":            47,
	"]":            48,
	"\\":           49,
	";":            51,
	"'":            52,
	"`":            53,
	",":            54,
	".":            55,
	"/":            56,
	"capslock":     57,
	"f1":           58,
	"f2":           59,
	"f3":           60,
	"f4":           61,
	"f5":           62,
	"f6":           63,
	"f7":           64,
	"f8":           65,
	"f9":           66,
	"f10":          67,
	"f11":          68,
	"f12":          69,
	"printscreen":  70,
	"scrolllock":   71,
	"pause":        72,
	"insert":       73,
	"home":         74,
	"pageup":       75,
	"delete":       76,
	"end":          77,
	"pagedown":     78,
	"right":        79,
	"left":         80,
	"down":         81,
	"up":           82,
	"num_lock":     83,
	"num_/":        84,
	"num_*":        85,
	"num_-":        86,
	"num_+":        87,
	"num_enter":    88,
	"num_1":        89,
	"num_2":        90,
	"num_3":        91,
	"num_4":        92,
	"num_5":        93,
	"num_6":        94,
	"num_7":        95,
	"num_8":        96,
	"num_9":        97,
	"num_0":        98,
	"num_.":        99,
	"application":  101,
	"control":      224,
	"shift":        225,
	"lshift":       225,
	"alt":          226,
	"windows":      227,
	"rightcontrol": 228,
	"rightshift":   229,
	"rightalt":     230,
	"rightwindows": 231,
}

type flyEngine struct {
	Dll    *syscall.LazyDLL
	Handle uintptr
	// proc
	Proc map[string]*syscall.LazyProc
}

var flye = &flyEngine{}

// OpenFly
func OpenFly(vid, pid int) (int, int, error) {
	// 引入dll
	flye.Dll = syscall.NewLazyDLL("fly.dll")
	err := flye.Dll.Load()
	if err != nil {
		return -1, -1, errors.New("Need fly.dll")
	}
	// init proc
	flye.Proc = map[string]*syscall.LazyProc{
		"M_Open":           flye.Dll.NewProc("M_Open"),
		"M_Open_VidPid":    flye.Dll.NewProc("M_Open_VidPid"),
		"M_ResolutionUsed": flye.Dll.NewProc("M_ResolutionUsed"),
		"M_ReleaseAllKey":  flye.Dll.NewProc("M_ReleaseAllKey"),
		"M_KeyDown":        flye.Dll.NewProc("M_KeyDown"),
		"M_KeyUp":          flye.Dll.NewProc("M_KeyUp"),
		"M_MoveTo3_D":      flye.Dll.NewProc("M_MoveTo3_D"),
		"M_MoveTo3":        flye.Dll.NewProc("M_MoveTo3"),
		"M_MouseWheel":     flye.Dll.NewProc("M_MouseWheel"),
		"M_LeftUp":         flye.Dll.NewProc("M_LeftUp"),
		"M_LeftDown":       flye.Dll.NewProc("M_LeftDown"),
		"M_RightUp":        flye.Dll.NewProc("M_RightUp"),
		"M_RightDown":      flye.Dll.NewProc("M_RightDown"),
	}

	// 打开端口
	var openDevice = func() (int, int) {
		// https://docs.microsoft.com/zh-cn/windows/win32/cimwin32prov/win32-diskdrive
		var storageinfo []struct {
			DeviceID string
		}
		err2 := wmi.Query("Select * from Win32_Keyboard ", &storageinfo)
		if err2 != nil {
			return -1, -1
		}
		r, _ := regexp.Compile("(?U)ID_(.*)&")
		for _, v := range storageinfo {
			idArr := r.FindAllStringSubmatch(v.DeviceID, -1)
			if len(idArr) == 2 {
				vidStr, pidStr := idArr[0][1], idArr[1][1]
				vid, _ := strconv.ParseInt("0x"+vidStr, 0, 64)
				pid, _ := strconv.ParseInt("0x"+pidStr, 0, 64)
				if vid > 0 && pid > 0 {
					flye.Handle, _, _ = flye.Proc["M_Open_VidPid"].Call(uintptr(vid), uintptr(pid))
					if handleOk() {
						return int(vid), int(pid)
					}
				}
			}
		}
		return -1, -1
	}
	flye.Handle, _, _ = flye.Proc["M_Open"].Call(uintptr(1))
	if !handleOk() {
		if vid > 0 {
			flye.Handle, _, _ = flye.Proc["M_Open_VidPid"].Call(uintptr(vid), uintptr(pid))
			if !handleOk() {
				vid, pid = openDevice()
			}
		} else {
			vid, pid = openDevice()
		}
	} else {
		vid, pid = 0, 0
	}
	if !handleOk() {
		return -1, -1, errors.New("Need flye.usb")
	}
	SwichEngine("fly")
	// set ResolutionUsed
	x, y := Screen.W, Screen.H
	flye.Proc["M_ResolutionUsed"].Call(flye.Handle, uintptr(x), uintptr(y))
	// release
	ReleaseFly()
	return vid, pid, nil
}

func handleOk() bool {
	if flye.Handle > 18446744073709 {
		return false
	}
	if flye.Handle < 1 {
		return false
	}
	return true
}

func ReleaseFly() {
	if flye.Handle < 1 {
		return
	}
	flye.Proc["M_ReleaseAllKey"].Call(flye.Handle)
}

func (f *flyEngine) keyDown(key string) {
	if f.Handle < 1 {
		return
	}
	mNum, has := flyKeyCode[key]
	if has == false {
		return
	}
	f.Proc["M_KeyDown"].Call(f.Handle, uintptr(mNum))
	return
}

func (f *flyEngine) keyUp(key string) {
	if f.Handle < 1 {
		return
	}
	mNum, has := flyKeyCode[key]
	if has == false {
		return
	}
	f.Proc["M_KeyUp"].Call(f.Handle, uintptr(mNum))
	return
}

func (f *flyEngine) moveTo(x int, y int) {
	if f.Handle < 1 {
		return
	}
	var p win.POINT
	for i := 0; i < 5; i++ {
		f.Proc["M_MoveTo3_D"].Call(f.Handle, uintptr(x), uintptr(y))
		Sleep(13 + rand.Intn(9))
		res := win.GetCursorPos(&p)
		if res && x == int(p.X) && y == int(p.Y) {
			return
		}
		Sleep(13 + rand.Intn(9))
	}
}

func (f *flyEngine) smoothMoveTo(x int, y int) {
	if f.Handle < 1 {
		return
	}
	var p win.POINT
	for i := 0; i < 5; i++ {
		f.Proc["M_MoveTo3"].Call(f.Handle, uintptr(x), uintptr(y))
		Sleep(13 + rand.Intn(9))
		res := win.GetCursorPos(&p)
		if res && x == int(p.X) && y == int(p.Y) {
			return
		}
		Sleep(13 + rand.Intn(9))
	}
}

func (f *flyEngine) mouseToggle(upDown, leftRight string) {
	if f.Handle < 1 {
		return
	}
	mBtn := "Down"
	if upDown == "up" {
		mBtn = "Up"
	}
	mKey := "M_Left"
	if leftRight == "right" {
		mKey = "M_Right"
	}
	f.Proc[mKey+mBtn].Call(f.Handle)
}

func (f *flyEngine) mouseWheel(dist string) {
	if f.Handle < 1 {
		return
	}

	d := -1
	if dist == "up" {
		d = 1
	}
	f.Proc["M_MouseWheel"].Call(f.Handle, uintptr(d))
}
