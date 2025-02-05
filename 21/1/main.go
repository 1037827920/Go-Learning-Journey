// Description: 通过&获取变量的地址，通过*来解引用
// &无法获取字面值的地址

package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)
}