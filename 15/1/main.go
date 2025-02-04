// Description: 动态地添加切片长度，通过append函数

package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")

	fmt.Println(dwarfs)
}