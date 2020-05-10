package kit

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

type Point struct {
	X, Y int
}

// MoveTo
func (p Point) MoveTo() {
	robotgo.Move(p.X, p.Y)
}

// GetColor get the point color
func (p Point) GetColor() string {
	return robotgo.PadHex(robotgo.GetPxColor(p.X, p.Y))
}

// DragTo start on a point, drag to another point.
func (p Point) DragTo(d Point, sleep int) {
	MoveTo(p.X, p.Y)
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("down")
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("down")
	Sleep(66 + rand.Intn(22))
	SmoothTo(d, sleep)
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("up")
	Sleep(66 + rand.Intn(22))
	robotgo.MouseToggle("up")
	Sleep(66 + rand.Intn(22))
}

// SmoothTo start on a point, to another point.
func (p Point) SmoothTo(d Point, sleep int) {
	tmp := p
	i := 0
	for {
		var m Point
		if d.X-tmp.X < 2 {
			m.X = d.X - tmp.X
		} else {
			m.X = rand.Intn(2)
		}
		if d.Y-tmp.Y < 2 {
			m.Y = d.Y - tmp.Y
		} else {
			m.Y = rand.Intn(2) + 1
		}
		tmp = tmp.Plus(m)
		tmp.MoveTo()
		if tmp == d {
			return
		}
		//
		if d.X-tmp.X < 20 && d.Y-tmp.Y < 20 {
			if i%2 == 0 {
				time.Sleep(time.Duration(sleep*2) * time.Nanosecond)
			}
		} else {
			if i < 55 {
				time.Sleep(time.Duration(sleep) * time.Nanosecond)
			} else if i%3 == 0 {
				time.Sleep(time.Duration(sleep) * time.Nanosecond)
			}
		}
		i++
	}
}

// Add
func (p Point) Add(w, h int) Point {
	return Point{p.X + w, p.Y + h}
}

// Square
func (p Point) Square(a int) Point {
	return Point{p.X + a, p.Y + a}
}

// Plus
func (p Point) Plus(a Point) Point {
	return Point{p.X + a.X, p.Y + a.Y}
}

// RightClick mouse right click
func (p Point) RightClick() {
	robotgo.Move(p.X, p.Y)
	Sleep(33 + rand.Intn(33))
	robotgo.MouseToggle("down", "right")
	Sleep(33 + rand.Intn(33))
	robotgo.MouseToggle("up", "right")
}

// LeftClick mouse left click
func (p Point) LeftClick() {
	p.MoveTo()
	robotgo.MouseToggle("down", "left")
	Sleep(55 + rand.Intn(10))
	robotgo.MouseToggle("up", "left")
}

// LeftDoubleClick mouse left double click
func (p Point) LeftDoubleClick() {
	p.MoveTo()
	robotgo.MouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("up", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("down", "left")
	Sleep(99 + rand.Intn(22))
	robotgo.MouseToggle("up", "left")
	Sleep(99 + rand.Intn(22))
}

func (p Point) Scroll(dist string) {
	if dist != "up" {
		dist = "down"
	}
	p.MoveTo()
	Sleep(88 + rand.Intn(22))
	robotgo.ScrollMouse(1, dist)
	Sleep(88 + rand.Intn(22))
}

func (p Point) String() string {
	return "P(" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + ")"
}

func (p Point) Area(w, h int) Area {
	return Area{p.X, p.Y, w, h}
}
