package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {

	f, _ := os.Create("imgstr/imgstr.go")
	defer f.Close()
	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "%v\n", "package imgstr")
	fmt.Fprintf(w, "%v\n", "var ImgStr map[string]robotgo.CBitmap")
	fmt.Fprintf(w, "%v\n", "func init() {")
	fmt.Fprintf(w, "%v\n", "ImgStr = map[string]robotgo.CBitmap{")

	path := "."
	filepath.Walk(path, func(imgsrc string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		bit := robotgo.OpenBitmap(imgsrc, 2)
		str := robotgo.TostringBitmap(bit)
		imgsrcStr := strings.Replace(imgsrc, "\\", "/", -1)
		bitmapStr := "robotgo.CBitmap(robotgo.BitmapStr(\"" + str + "\"))"
		fmt.Fprintf(w, "%v\n", "\""+imgsrcStr+"\":"+bitmapStr+",")
		return nil
	})
	fmt.Fprintf(w, "%v\n", "}}")
	w.Flush()
}
