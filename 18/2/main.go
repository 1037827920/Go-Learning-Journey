// Description: 构造函数

package main

import "fmt"

type coordinate struct {
	d, m, s float64 // 度, 分, 秒
	h       rune    // 南纬(S)或北纬(N), 东经(E)或西经(W)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	loc := newLocation(lat, long)
	fmt.Println(loc)
}