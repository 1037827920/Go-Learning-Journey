// Description: 声明可变参数的函数

package main

import "fmt"

// 通过...标记，将worlds参数标记为可变参数
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	planets := []string{"Venus", "Mars", "Jupiter"}
	// 通过...将切片展开
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}