package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/vcgo/kit"
)

func main() {
	// params
	c := flag.Uint64("c", 0xff00ff, "start hex color")
	w := flag.Int("w", -1, "color point add w")
	h := flag.Int("h", -1, "color point add p")
	flag.Parse()
	if len(os.Args) < 7 {
		flag.PrintDefaults()
		return
	}
	// count
	bmp := kit.Screen.Capture()
	defer bmp.Free()
	p, err := bmp.FindColor(robotgo.UintToHex(uint32(*c)), 0.0)
	if err != nil {
		fmt.Println("hexmatrix can't find color")
		return
	}
	// output & save
	for j := 0; j < *h; j++ {
		for i := 0; i < *w; i++ {
			fmt.Printf("0x%v ", p.Add(i, j).GetColor())
		}
		fmt.Println("")
	}
	area := p.Area(*w, *h)
	area.Test("hexmatrix", "./tests")
}
