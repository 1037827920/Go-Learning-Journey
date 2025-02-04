// Description: 使用make对slice进行预分配

package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}


func main() {
	dwarfs := make([]string, 0, 10)

	dump("dwarfs", dwarfs)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	dump("dwarfs", dwarfs)
}