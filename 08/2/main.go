// Description: 类型别名

package main

import "fmt"

func main() {
	type byte = uint8
	type rune = int32

	var a byte = 'a'
	var b rune = '中'

	fmt.Printf("%c %c\n", a, b)
}