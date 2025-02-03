// Description: 匿名函数

package main

import "fmt"

var f = func() {
	fmt.Println("Hello, World!")
}

func main() {
	f()

	// 使用短声明创建匿名函数
	f1 := func(message string) {
		fmt.Println(message)
	}
	f1("Hello, World!")

	// 声明完匿名参数后立即执行
	func() {
		fmt.Println("Hello, World!")
	}()
}