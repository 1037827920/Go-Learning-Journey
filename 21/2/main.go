// Description: 指针类型

package main

import "fmt"

func main() {
	answer := 42
	address := &answer

	fmt.Printf("address is a %T\n", address)
}