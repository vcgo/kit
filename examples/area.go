package main

import (
	"fmt"

	"github.com/vcgo/kit"
)

func main() {
	areas := kit.Screen.Splice(3, 2)
	fmt.Println("...", areas)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			area := areas[i][j]
			area.Test("caches")
			fmt.Println("...", i, j, area)
			kit.Sleep(2000)
		}
	}
}
