package kit

import "github.com/go-vgo/robotgo"

// Bitmap is a pix-map type.
type Bitmap robotgo.Bitmap

func (area Area) Capture() Bitmap {
	cbitmap := robotgo.CaptureScreen(area.X, area.Y, area.W, area.H)
	return Bitmap(robotgo.ToBitmap(cbitmap))
}

func (bm Bitmap) ToString() string {
	bitmap := robotgo.Bitmap(bm)
	return robotgo.TostringBitmap(robotgo.ToCBitmap(bitmap))
}

func (bm Bitmap) ToBytes() []byte {
	bitmap := robotgo.Bitmap(bm)
	return robotgo.ToBitmapBytes(robotgo.ToCBitmap(bitmap))
}

func (bm Bitmap) SavePng(pngName string) error {
	cbitmap := robotgo.ToCBitmap(robotgo.Bitmap(bm))
	robotgo.SaveBitmap(cbitmap, pngName)
	return nil
}

func (bm Bitmap) Free() {
	cbitmap := robotgo.ToCBitmap(robotgo.Bitmap(bm))
	robotgo.FreeBitmap(cbitmap)
}

// // StringToBitmap trans string to Bitmap
// func StringToBitmap(str string) Bitmap {
// 	return Bitmap(robotgo.ToBitmap(robotgo.BitmapStr(str)))
// }
