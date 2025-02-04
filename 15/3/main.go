// Description: 三索引切分操作，第三个索引用来限制新建切片的容量

package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)
}