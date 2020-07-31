package kit

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/go-vgo/robotgo"
)

// go test -run TestFindHexMatrix
func TestFindHexMatrix(t *testing.T) {
	Screen := Area{0, 0, 800, 600}
	var color1 uint32 = 0xf6f8fe
	var color2 uint32 = 0x4e6ef2
	hm := HexMatrix{
		color1,
		[]HexPoint{
			{Point{0, 1}, color1},
			{Point{0, 2}, color1},
			{Point{0, 3}, color1},
			{Point{0, 4}, color1},
			{Point{0, 5}, color1},
			{Point{0, 6}, color1},
			{Point{0, 7}, color1},
			{Point{0, 8}, color1},
		},
		[]HexArea{
			{Area{-10, -10, 8, 8}, color2, 64},
			{Area{10, -10, 8, 8}, color2, 64},
			{Area{-10, 10, 8, 8}, color2, 32},
			{Area{10, 10, 8, 8}, color2, 32},
		},
	}
	startHm := time.Now()
	for i := 0; i < 10; i++ {
		start := time.Now()
		bmp := Screen.Capture()
		p, err := bmp.FindHexMatrix(hm)
		fmt.Println(time.Since(start), p, err, runtime.NumGoroutine())
		bmp.Free()
	}
	fmt.Println("========FindHexMatrix", time.Since(startHm))

	startHmGo := time.Now()
	for i := 0; i < 10; i++ {
		bmp := Screen.Capture()
		start := time.Now()
		p, err := bmp.FindHexMatrixGo(hm)
		fmt.Println(time.Since(start), p, err, runtime.NumGoroutine())
		defer bmp.Free()
	}
	fmt.Println("========FindHexMatrixGo", time.Since(startHmGo))

	findBmp := NewBitmap(robotgo.ToBitmap(robotgo.BitmapStr("b17,22,eNr7lOf3aQSj/6SAPzcvYJrwfXH/vx/fsErhQoNcy78vn4Ce/XV8F/Fa/j5/BFQP1PXvw5uvU2qIdxhQMVALQesw/QKxDmgvSd7HHybDLypJ0kIjBABVnL4p")))
	startFp := time.Now()
	for i := 0; i < 10; i++ {
		start := time.Now()
		p, err := Screen.FindPic(findBmp, 0.01)
		fmt.Println(time.Since(start), p, err, runtime.NumGoroutine())
	}
	fmt.Println("========FindPic", time.Since(startFp), color2)
}

func TestFindHexMatrixGo(t *testing.T) {
	var color1 uint32 = 0xf6f8fe
	var color2 uint32 = 0x4e6ef2
	hm := HexMatrix{
		color1,
		[]HexPoint{
			{Point{0, 1}, color1},
			{Point{0, 2}, color1},
			{Point{0, 3}, color1},
			{Point{0, 4}, color1},
			{Point{0, 5}, color1},
			{Point{0, 6}, color1},
			{Point{0, 7}, color1},
			{Point{0, 8}, color1},
		},
		[]HexArea{
			{Area{-10, -10, 8, 8}, color2, 64},
			{Area{10, -10, 8, 8}, color2, 64},
			{Area{-10, 10, 8, 8}, color2, 32},
			{Area{10, 10, 8, 8}, color2, 32},
		},
	}
	bmp := Screen.Capture()
	start := time.Now()
	p, err := bmp.FindHexMatrixGo(hm)
	fmt.Println(time.Since(start), p, err, runtime.NumGoroutine())
	defer bmp.Free()
	time.Sleep(time.Second * 5)
}
