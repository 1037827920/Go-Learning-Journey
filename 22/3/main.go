// Description: 当变量被赋值为函数类型时，它的默认值是nil

package main

import "fmt"

func main() {
	var fn func(a int, b int) int
	fmt.Println(fn == nil)
}