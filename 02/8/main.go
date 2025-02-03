// Description: 练习题，飞船距离目的地距离为56,000,000公里，
// 飞船要在28天内到达，行进速度为多少千米/小时

package main

import "fmt"

func main() {
	const distance = 56_000_000
	const days = 28
	const hours = 24
	const speed = distance / (days * hours)
	fmt.Printf("飞船的行进速度为: %v km/h \n", speed)
}
