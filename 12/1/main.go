// Description: 将函数赋值给变量

package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	// 将函数赋值给变量，并不执行函数
	sensor := fakeSensor
	// 只有遇到小括号才执行函数
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}