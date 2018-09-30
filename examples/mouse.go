package main

import (
	"github.com/vcgo/kit"
)

func main() {
	kit.Point{403, 372}.DragTo(kit.Point{565, 406})
	kit.Sleep(999)
	kit.LeftDoubleClick()
}
