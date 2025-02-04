// Description: 不通过切分数组来获得Slice，可以直接声明

package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// 通过切分数组来获取slice
	dwarfSlice := dwarfArray[:]

	// 直接声明slice
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfSlice)
	fmt.Println(dwarfs)

	fmt.Printf("dwarfArray: %T\n", dwarfArray)
	fmt.Printf("dwarfSlice: %T\n", dwarfs)
}