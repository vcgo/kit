package kit

import "github.com/wilon/robotg"

// Bitmap is a pix-map type.
type Bitmap robotg.Bitmap

func (area Area) Capture() Bitmap {
	cbitmap := robotg.CaptureScreen(area.X, area.Y, area.W, area.H)
	return Bitmap(robotg.ToBitmap(cbitmap))
}

func (bm Bitmap) ToString() string {
	bitmap := robotg.Bitmap(bm)
	return robotg.TostringBitmap(robotg.ToCBitmap(bitmap))
}

// func (bm Bitmap) ToBytes() []byte {
// 	bitmap := robotg.Bitmap(bm)
// 	return robotg.ToBitmapBytes(robotg.ToCBitmap(bitmap))
// }

func (bm Bitmap) SavePng(pngName string) error {
	cbitmap := robotg.ToCBitmap(robotg.Bitmap(bm))
	robotg.SaveBitmap(cbitmap, pngName)
	return nil
}

func (bm Bitmap) Free() {
	cbitmap := robotg.ToCBitmap(robotg.Bitmap(bm))
	robotg.FreeBitmap(cbitmap)
}

// // StringToBitmap trans string to Bitmap
// func StringToBitmap(str string) Bitmap {
// 	return Bitmap(robotg.ToBitmap(robotg.BitmapStr(str)))
// }
