package kit

import (
	"errors"
	"image"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/oliamb/cutter"
)

// Bitmap is a pix-map type.
// Cover tips:
// 		To robotgo CBitmap: `cbitmap := robotgo.ToCBitmap(robotgo.Bitmap(bm))`
type Bitmap struct {
	robotgo.Bitmap
	adbfile string
}

func NewBitmap(bit robotgo.Bitmap, args ...string) Bitmap {
	file := ""
	if len(args) > 0 {
		file = args[0]
	}
	return Bitmap{bit, file}
}

func (area Area) Capture() Bitmap {
	if engine == "adb" {
		bitmap, _ := area.adbCapture()
		return bitmap
	} else {
		cbitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
		return NewBitmap(robotgo.ToBitmap(cbitmap))
	}
}

func (wherebmp Bitmap) FindBitmap(findbmp Bitmap, tolerance float64) (Point, error) {
	findcbmp := robotgo.ToCBitmap(findbmp.Bitmap)
	wherecbmp := robotgo.ToCBitmap(wherebmp.Bitmap)
	x, y := robotgo.FindBitmap(findcbmp, wherecbmp, tolerance)
	if x > 0 || y > 0 {
		return Point{x, y}, nil
	} else {
		return Point{-1, -1}, errors.New("Cant find bitmap")
	}
}

func (bm Bitmap) ToString() string {
	return robotgo.TostringBitmap(robotgo.ToCBitmap(bm.Bitmap))
}

// func (bm Bitmap) ToBytes() []byte {
// 	bitmap := robotgo.Bitmap(bm)
// 	return robotgo.ToBitmapBytes(robotgo.ToCBitmap(bitmap))
// }

func (bm Bitmap) SavePng(pngName string) error {
	cbitmap := robotgo.ToCBitmap(bm.Bitmap)
	robotgo.SaveBitmap(cbitmap, pngName)
	return nil
}

func (bm Bitmap) Free() {
	cbitmap := robotgo.ToCBitmap(bm.Bitmap)
	robotgo.FreeBitmap(cbitmap)
	if engine == "adb" && bm.adbfile != "" {
		os.Remove(bm.adbfile)
	}
}

// // StringToBitmap trans string to Bitmap
// func StringToBitmap(str string) Bitmap {
// 	return Bitmap(robotgo.ToBitmap(robotgo.BitmapStr(str)))
// }

func (a Area) adbCapture() (Bitmap, error) {
	// Get .png output.
	pngName := string(time.Now().Format("2006_01_02.15_04_05"))
	pngName += strconv.Itoa(rand.Intn(99999)) + "_adbadb.png"
	file := path.Join(os.TempDir(), pngName)
	shell := "screencap -p"
	output, _ := device.RunCommand(shell)
	lines := strings.Split(output, "\n")
	pngStr := strings.Join(lines[2:], "\n")
	pngReader := strings.NewReader(pngStr)
	out, _ := os.Create(file)
	defer out.Close()
	// Need clip ?
	if a != Screen {
		pngImg, fm, err := image.Decode(pngReader)
		if err != nil || fm != "png" {
			return Bitmap{}, errors.New("Capture png error.")
		}
		croppedImg, err := cutter.Crop(pngImg, cutter.Config{
			Width:  a.W,
			Height: a.H,
			Anchor: image.Point{a.X, a.Y},
			Mode:   cutter.TopLeft,
		})
		png.Encode(out, croppedImg)
	} else {
		io.Copy(out, pngReader)
	}
	// Capture
	cbitmap := robotgo.OpenBitmap(file)
	bit := NewBitmap(robotgo.ToBitmap(cbitmap), file)
	return bit, nil
}
