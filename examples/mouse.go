package main

import (
	"github.com/vcgo/kit"
)

func main() {
	kit.Point{200, 33}.DragTo(kit.Point{500, 500})
	kit.Sleep(999)
	kit.LeftDoubleClick()
}
