package main

import (
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wilon/robotg"
)

func main() {

	code := ""

	code += "package imgstr\n"
	code += "import (\n"
	code += " \"github.com/wilon/robotg\"\n"
	code += " \"github.com/vcgo/kit\"\n"
	code += ")\n"
	code += "var ImgStr map[string]kit.Bitmap\n"
	code += "func init() {\n"
	code += "ImgStr = map[string]kit.Bitmap{\n"

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

		bit := robotg.OpenBitmap(imgsrc, imgType)
		str := robotg.TostringBitmap(bit)
		imgsrcStr := strings.Replace(imgsrc, "\\", "/", -1)
		bitmapStr := "kit.Bitmap(robotg.ToBitmap(robotg.BitmapStr(\"" + str + "\")))"
		code += "\"" + imgsrcStr + "\":" + bitmapStr + ",\n"
		return nil
	})

	code += "}}"
	res, err := format.Source([]byte(code))
	if err != nil {
		panic(err)
	}
	err2 := ioutil.WriteFile("./imgstr.go", res, 0755)
	if err2 != nil {
		panic(err)
	}
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
