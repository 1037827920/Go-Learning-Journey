// Description: 将切片作为参数传递给函数，会对原始切片进行修改。

package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	planets := []string{" Venus ", " Earth ", " Mars "}
	hyperspace(planets)

	fmt.Println(strings.Join(planets, ""))
}