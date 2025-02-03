// Description: 声明方法，可以将方法与类型绑定在一起，但不能绑定
// 编译器如int，float64等内建类型

package main

import "fmt"

type celsius float64
type kelvin float64

// kelvintoCelsius函数将kelvin类型转换为celsius类型
func kelvintoCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// 将上面这个函数声明为kelvin类型的方法，方法名为celsius()
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	c1 := kelvintoCelsius(k)
	c2 := k.celsius()
	fmt.Println(c1, c2)
}