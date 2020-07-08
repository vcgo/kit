package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {

	code1 := `package imgstr

import (
	"github.com/go-vgo/robotgo"
	"github.com/vcgo/kit"
)

var strmap map[string]string

func init() {
	strmap = map[string]string{
`
	code2 := ""
	code3 := `
	}
}

func Get(str string) kit.Bitmap {
	v, ok := strmap[str]
	if ok {
		return kit.NewBitmap(robotgo.ToBitmap(robotgo.BitmapStr(v)))
	} else {
		return kit.Bitmap{}
	}
}
`
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
			imgType = 1
		case ".bmp":
			imgType = 2
		default:
			return nil
		}
		if len(os.Args) == 2 && os.Args[1] == "-v" {
			fmt.Println("Add", imgsrc)
		}
		bit := robotgo.OpenBitmap(imgsrc, imgType)
		str := robotgo.TostringBitmap(bit)
		imgsrcStr := strings.Replace(imgsrc, "\\", "/", -1)
		bitmapStr := "\"" + str + "\""
		code2 += "\"" + imgsrcStr + "\":" + bitmapStr + ","
		return nil
	})

	res, err := format.Source([]byte(code1 + code2 + code3))
	if err != nil {
		fmt.Println("transimg format.Source error!")
		panic(err)
	}
	err2 := ioutil.WriteFile("./imgstr.go", res, 0755)
	if err2 != nil {
		fmt.Println("transimg ioutil.WriteFile error!")
		panic(err)
	}
	fmt.Println("transimg success!")
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
