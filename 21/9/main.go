// Description: 指向slice的显式指针
// 这个显式指针的唯一作用就是修改slice本身：slice的长度、容量以及起始偏移量

package main

import "fmt"

func reclassify(planets *[]string) {
	// 将长度修改为8
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)

	fmt.Println(planets)
}