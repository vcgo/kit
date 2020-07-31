package kit

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
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

func (bmp Bitmap) FindBitmap(findbmp Bitmap, tolerance float64) (Point, error) {
	findcbmp := robotgo.ToCBitmap(findbmp.Bitmap)
	wherecbmp := robotgo.ToCBitmap(bmp.Bitmap)
	x, y := robotgo.FindBitmap(findcbmp, wherecbmp, tolerance)
	if x > 0 || y > 0 {
		return Point{x, y}, nil
	} else {
		return Point{-1, -1}, errors.New("Cant find bitmap")
	}
}

// HexPoint
// HexMatrix find this Point Hex
type HexPoint struct {
	Point Point
	Hex   uint32
}

// HexArea
// HexMatrix find this Area count Hex >= Count
type HexArea struct {
	Area  Area
	Hex   uint32
	Count int
}

type HexMatrix struct {
	Hex    uint32
	Points []HexPoint
	Areas  []HexArea
}

func (bmp Bitmap) FindHexMatrix(hm HexMatrix) (Point, error) {
	wherecbmp := robotgo.ToCBitmap(bmp.Bitmap)
	w, h := bmp.Bitmap.Width, bmp.Bitmap.Height
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			hex := robotgo.CHex(robotgo.GetColor(wherecbmp, x, y))
			if hex == robotgo.UintToHex(hm.Hex) {
				match := true
				// find points
				for _, p := range hm.Points {
					m, n := x+p.Point.X, y+p.Point.Y
					hex := robotgo.CHex(robotgo.GetColor(wherecbmp, m, n))
					if hex != robotgo.UintToHex(p.Hex) {
						match = false
						break
					}
				}
				// count areas
				for _, a := range hm.Areas {
					count := 0
					for m := x + a.Area.X; m < x+a.Area.X+a.Area.W; m++ {
						for n := y + a.Area.Y; n < y+a.Area.Y+a.Area.H; n++ {
							hex := robotgo.CHex(robotgo.GetColor(wherecbmp, m, n))
							if hex == robotgo.UintToHex(a.Hex) {
								count++
								if count >= a.Count {
									break
								}
							}
						}
						if count >= a.Count {
							break
						}
					}
					if count < a.Count {
						match = false
						break
					}
				}
				// res
				if match {
					return Point{x, y}, nil
				}
			}
		}
	}
	return Point{-1, -1}, errors.New("Cant find HexMatrix")
}

func (bmp Bitmap) FindHexMatrixGo(hm HexMatrix) (Point, error) {
	wherecbmp := robotgo.ToCBitmap(bmp.Bitmap)
	w, h := bmp.Bitmap.Width, bmp.Bitmap.Height
	wg := &sync.WaitGroup{}
	resCh := make(chan Point)
	doneCh := make(chan bool)
	for x := 0; x < w; x++ {
		wg.Add(1)
		go func(x int, resCh chan Point) {
			defer wg.Done()
			for y := 0; y < h; y++ {
				hex := robotgo.CHex(robotgo.GetColor(wherecbmp, x, y))
				if hex == robotgo.UintToHex(hm.Hex) {
					match := true
					// find points
					for _, p := range hm.Points {
						m, n := x+p.Point.X, y+p.Point.Y
						hex := robotgo.CHex(robotgo.GetColor(wherecbmp, m, n))
						if hex != robotgo.UintToHex(p.Hex) {
							match = false
							break
						}
					}
					// count areas
					for _, a := range hm.Areas {
						count := 0
						for m := x + a.Area.X; m < x+a.Area.X+a.Area.W; m++ {
							for n := y + a.Area.Y; n < y+a.Area.Y+a.Area.H; n++ {
								hex := robotgo.CHex(robotgo.GetColor(wherecbmp, m, n))
								if hex == robotgo.UintToHex(a.Hex) {
									count++
									if count >= a.Count {
										break
									}
								}
							}
							if count >= a.Count {
								break
							}
						}
						if count < a.Count {
							match = false
							break
						}
					}
					// res
					if match {
						resCh <- Point{x, y}
						for {
							select {
							case <-doneCh:
								return
							}
						}
					}
				}
			}
		}(x, resCh)
	}
	go func() {
		wg.Wait()
		fmt.Println("wg.Wait finished")
		resCh <- Point{-1, -1}
	}()
	for {
		select {
		case p := <-resCh:
			go func() {
				for i := 0; i < 99; i++ {
					doneCh <- true
				}
			}()
			if p.X < 0 {
				return Point{-1, -1}, errors.New("Cant find HexMatrix")
			}
			return p, nil
		}
	}
}

func (bmp Bitmap) FindColor(color robotgo.CHex, tolerance float64) (Point, error) {
	var x, y int
	wherecbmp := robotgo.ToCBitmap(bmp.Bitmap)
	x, y = robotgo.FindColor(color, wherecbmp, tolerance)
	if x > 0 || y > 0 {
		return Point{x, y}, nil
	} else {
		return Point{x, y}, errors.New("Cant find color")
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
	if len(lines) <= 2 {
		return Bitmap{}, errors.New("Capture png error.")
	}
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
