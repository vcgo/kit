package kit

import (
	"math/rand"
)

// MoveTo mouse move to x, y
func MoveTo(x int, y int) {
	for i := 0; i < 55; i++ {
		SerialWrite([]byte{0xAA, 0x61, 0x04, 0x00, 127, 127, 0x00})
	}

	// 距离
	xDist, yDist := int(Screen.W)-x, int(Screen.H)-y

	// 哪个长做master
	direction, master, slave := "x", xDist, yDist
	if xDist < yDist {
		direction = "y"
		master, slave = yDist, xDist
	}

	// 每步125
	step := 125
	masterStep, slaveStep := step, step // 初始step
	xStep, yStep := 0, 0                // 结果
	times := int(master / step)
	for i := times; i >= 0; i-- {
		// master最后一步
		if i == 0 {
			masterStep = master % step
		}
		// slave结束的早
		if slave <= slaveStep {
			slaveStep = slave
		}
		slave -= slaveStep

		if direction == "x" {
			xStep = masterStep
			yStep = slaveStep
		} else {
			xStep = slaveStep
			yStep = masterStep
		}

		SerialWrite([]byte{0xAA, 0x61, 0x04, 0x00, getMoveByte(xStep), getMoveByte(yStep), 0x00})
	}
}

func getMoveByte(step int) byte {
	var res int
	if step > 0 {
		res = 256 - step
	} else if step == 0 {
		res = 0
	} else {
		res = 256 + step
	}
	return byte(res)
}

// LeftClick mouse left click
func LeftClick() {
	SerialWrite([]byte{0xAA, 0x61, 0x04, 0x01, 0, 0, 0})
	Sleep(55 + rand.Intn(20))
	SerialWrite([]byte{0xAA, 0x61, 0x04, 0x00, 0, 0, 0})
}

// RightClick mouse right click
func RightClick() {
	SerialWrite([]byte{0xAA, 0x61, 0x04, 0x02, 0, 0, 0})
	Sleep(55 + rand.Intn(20))
	SerialWrite([]byte{0xAA, 0x61, 0x04, 0, 0, 0, 0})
}

// LeftDoubleClick mouse left double click
func LeftDoubleClick() {
	LeftClick()
	Sleep(55 + rand.Intn(20))
	LeftClick()
}

// MoveClick mouse move to and left click
func MoveClick(x int, y int) {
	MoveTo(x, y)
	Sleep(88 + rand.Intn(20))
	LeftClick()
}

// MoveDoubleClick mouse move to and double click
func MoveDoubleClick(x int, y int) {
	MoveTo(x, y)
	Sleep(88 + rand.Intn(20))
	LeftDoubleClick()
}
