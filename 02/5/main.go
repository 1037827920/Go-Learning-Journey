// Description: 变量声明的三种方式

package main

import "fmt"

func main() {
	const distance = 56000000
	const speed = 100800

	const (
		distance2 = 56000000
		speed2    = 100800
	)

	const distance3, speed3 = 56000000, 100800

	const hoursePerDay, minutesPerHour = 24, 60

	fmt.Println(distance, speed)
	fmt.Println(distance2, speed2)
	fmt.Println(distance3, speed3)
	fmt.Println(hoursePerDay, minutesPerHour)
}
