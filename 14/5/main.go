// Description: 带有方法的切片
// 在go中可以将slice或数组作为底层类型，然后绑定其他方法

package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Sort() {
	
}

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}