// Description: 声明函数类型

package main

import (
	"fmt"
	"time"
	"math/rand"
)

type kelvin float64
// 声明函数类型
type sensor func() kelvin

func measureTemperature(samples int, sensor sensor) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor)
}
