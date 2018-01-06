# A kit.

## Install

only support .bmp

```bash
go get github.com/vcgo/kit
go get github.com/vcgo/kit/transimg

```

## Doc

```go
type Area struct{ ... }
var Logger log.Logger
func Exit(args ...interface{})
func FindColor(area Area, color robotgo.CHex) (int, int, error)
func FindPic(area Area, imgbitmap robotgo.CBitmap, tolerance float32) (int, int, error)
func KeyDown(key string)
func KeyPress(key string)
func KeyUp(key string)
func LeftClick()
func Log(desc string, args ...interface{})
func MoveTo(x int, y int)
func Sleep(x int)
```