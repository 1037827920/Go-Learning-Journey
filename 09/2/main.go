// Description: 使用strconv包进行类型转换

package main

import (
	"fmt"
	"strconv"
)

func main() {
	count_down := 10
	str := "Launch in T minus " + strconv.Itoa(count_down) + " seconds."
	fmt.Println(str)
}