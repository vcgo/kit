package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {

	// if mkdirs("imgstr") == false {
	// 	fmt.Println("mkdir imgstr failed!")
	// 	return
	// }

	f, _ := os.Create("imgstr.go")
	defer f.Close()
	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "%v\n", "package imgstr")
	fmt.Fprintf(w, "%v\n", "var ImgStr map[string]robotgo.CBitmap")
	fmt.Fprintf(w, "%v\n", "func init() {")
	fmt.Fprintf(w, "%v\n", "ImgStr = map[string]robotgo.CBitmap{")

	imgpath := "."
	filepath.Walk(imgpath, func(imgsrc string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		var imgType int
		switch path.Ext(imgsrc) {
		case ".png":
			// imgType = 0
			return nil
		case ".bmp":
			imgType = 2
		default:
			return nil
		}

		bit := robotgo.OpenBitmap(imgsrc, imgType)
		str := robotgo.TostringBitmap(bit)
		imgsrcStr := strings.Replace(imgsrc, "\\", "/", -1)
		bitmapStr := "robotgo.CBitmap(robotgo.BitmapStr(\"" + str + "\"))"
		fmt.Fprintf(w, "%v\n", "\""+imgsrcStr+"\":"+bitmapStr+",")
		return nil
	})
	fmt.Fprintf(w, "%v\n", "}}")
	w.Flush()
}

func mkdirs(imgpath string) bool {

	_, err := os.Stat(imgpath)
	if err == nil {
		return true
	}
	errm := os.MkdirAll(imgpath, 0755)
	if errm == nil {
		return true
	}
	return false

}
