// Description: 通过给slice中的所有行星加上New前缀
// 必须使用planets类型，并为之实现乡音给的terraform方法

package main

import "fmt"

type planets []string

func (p planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	planets := planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	planets.terraform()

	fmt.Println(planets)
}