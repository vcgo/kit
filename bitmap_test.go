package kit

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	bitmapStr := Area{10, 10, 6, 6}.Capture().ToString()
	fmt.Println("bitmapStr", bitmapStr)
}

// func TestToBytes(t *testing.T) {
// 	bitmapBytes := Area{10, 10, 6, 6}.Capture().ToBytes()
// 	fmt.Println("bitmapBytes", bitmapBytes)
// }
