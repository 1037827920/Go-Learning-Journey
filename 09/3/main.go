// Description: 错误处理初步接触

package main

import (
	"fmt"
	"strconv"
)

func main() {
	count_down, err := strconv.Atoi("10")
	if err != nil {
		// 错误处理
	}
	fmt.Println(count_down)
}