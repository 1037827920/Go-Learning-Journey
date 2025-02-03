// Description: 将布尔类型转换为整型，true转换为1，false转换为0

package main

import "fmt"

func main() {
	b := true

	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Println(i)
}