// Description: 组合结构体，更加灵活

package main

import (
	"fmt"
)

type report struct {
	sol         int
	temperature temperature
	location    location
}

// 这里是转发了temperature的average方法
func (r report) average() celsius {
	return r.temperature.average()
}

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
	fmt.Println(t.average())
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Println(report.average())

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %v C\n", report.temperature.high)
}
