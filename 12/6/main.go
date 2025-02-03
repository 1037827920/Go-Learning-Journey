// Description: 闭包的使用

package main

import "fmt"

type kelvin float64

func main() {
	var k kelvin = 294.0

	// 闭包可以捕获外部变量的值
	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor())

	k++
	fmt.Println(sensor())
}