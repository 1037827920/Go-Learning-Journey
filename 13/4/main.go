// Description: 将数组作为参数传递给函数，是传递的数组的副本，而不是原数组
// 所以函数一般使用slice作为参数，而不是数组

package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terraform(planets)
	fmt.Println(planets)
}