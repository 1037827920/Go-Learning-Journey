// Description: 转发方法，通过struct嵌入来实现方法的自动转发

package main

import (
	"fmt"
)

// struct嵌入，只给出类型，不给出字段名
// 这样temperature的方法可以直接给report使用，不用想上个例子那样写出来
type report struct {
	sol
	temperature
	location
}

type sol int

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type celsius float64

type location struct {
	lat, long float64
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	// 直接通过report调用temperature的方法
	fmt.Println(report.average())

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %v C\n", report.temperature.high)
}
