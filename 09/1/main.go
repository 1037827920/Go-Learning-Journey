// Description: 类型转换

package main

import (
	"fmt"
	"math"
)

func main() {
	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println(v8)
	}
}