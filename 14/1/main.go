// Description: Slice

package main

import "fmt"

func main() {
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Slice的默认索引
	// 前面如果忽略就默认从0开始
	// 后面如果忽略就默认到最后一个元素
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	// Slice索引不能为负数

	fmt.Println(terrestrial, gasGiants, iceGiants)
}
