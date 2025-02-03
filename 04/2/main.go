// Description: 短声明，等价于用var声明

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 在for循环中使用短声明
	for count:=0; count<10; count++ {
		fmt.Println(count)
	}

	// 在if中使用短声明
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	// 在switch中使用短声明
	switch num :=rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}