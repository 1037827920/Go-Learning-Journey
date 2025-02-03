// Description: string类型

package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	// string类型是不修改的
	// message[5] = 'd'
}