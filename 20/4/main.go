// Description: Go语言的接口都是隐式实现的

package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// func stardata(t time.Time) float64 {
// 	doy := float64(t.YearDay())
// 	h := float64(t.Hour()) / 24.0
// 	return 1000.0 + doy + h
// }

// time.Time实现了上面定义的dater接口，所以这里可以改成下面的写法
func stardata(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000.0 + doy + h
}

func main() {
	day := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardata(day))
}