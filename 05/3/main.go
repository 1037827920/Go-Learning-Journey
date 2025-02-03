// Description: 如何比较浮点类型

package main

import (
	"fmt"
	"math"
)

func main() {
	piggy_bank := 0.1
	piggy_bank += 0.2

	// 直接进行比较是错误的
	fmt.Println(piggy_bank == 0.3)

	// 正确地比较
	fmt.Println(math.Abs(piggy_bank-0.3) < 0.0001)
}