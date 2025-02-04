// Description: 用map实现set，set不存在重复元素

package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("Set member")
	}

	fmt.Println(set)

	// map遍历是无序的
	// 顺序输出map中的元素
	// 将map中的元素放到unique slice中
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}