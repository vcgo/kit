package main

import (
	"math/rand"

	"github.com/vcgo/kit"
)

// 安卓手机
// 打开画图，写一个口字
func main() {
	kit.Fmt("adb test")
	err := kit.OpenAdbDevice()
	if err != nil {
		panic(err)
	}
	kit.Sleep(1000)

	// 选色1
	kit.Point{880, 1928}.LeftClick()
	kit.Sleep(1000)
	// 滑动【口】第一笔
	kit.Point{300, 500}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("down", "left")
	kit.Sleep(555 + rand.Intn(99))
	kit.Point{300, 1000}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("up", "left")
	kit.Sleep(1000)
	// 选色2
	kit.Point{456, 1928}.LeftClick()
	kit.Sleep(1000)
	// 滑动【口】第二笔
	kit.Point{300, 500}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("down", "left")
	kit.Sleep(555 + rand.Intn(99))
	kit.Point{800, 500}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.Point{800, 1000}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("up", "left")
	kit.Sleep(1000)
	// 选色3
	kit.Point{25, 1928}.LeftClick()
	kit.Sleep(1000)
	// 滑动【口】第三笔
	kit.Point{300, 1000}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("down", "left")
	kit.Sleep(555 + rand.Intn(99))
	kit.Point{800, 1000}.MoveTo()
	kit.Sleep(555 + rand.Intn(99))
	kit.MouseToggle("up", "left")
	kit.Sleep(1000)
}
