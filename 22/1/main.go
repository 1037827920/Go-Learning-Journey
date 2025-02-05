// Description: nil

package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	// 下面的代码会发生panic
	// fmt.Println(*nowhere)
}