package kit

import (
	"fmt"
	"testing"
	"time"
)

// go test -run TestFindHexMatrix
func TestFindHexMatrix(t *testing.T) {
	start := time.Now()
	hm := HexMatrix{
		0x909497,
		[]HexDeviation{
			{0, 1, 0x909497},
			{0, 2, 0x909497},
			{0, 3, 0x909497},
			{0, 4, 0x909497},
			{0, 5, 0x909497},
			{0, 6, 0x909497},
			{0, 7, 0x909497},
			{0, 8, 0x909497},
			{0, 9, 0x909497},
			{0, 10, 0x909497},
		},
	}
	area := Area{0, 0, 800, 600}
	bmp := area.Capture()
	defer bmp.Free()
	p, err := bmp.FindHexMatrix(hm)
	fmt.Println(time.Since(start), p, err)
}
