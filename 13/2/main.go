// Description: 使用复合字面值初始化数组

package main

import "fmt"

func main() {
	// 通过...让编译器来推断数组的长度
	dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfs)
}