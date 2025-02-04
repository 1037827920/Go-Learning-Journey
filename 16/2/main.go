// Description: int、float64等类型在赋值给新变量或传递至函数时都会复制一个副本
// 但map不会

package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars": "Sector ZZ9",
	}	

	planetMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetMarkII)

	delete(planets, "Earth")
	fmt.Println(planets)
	fmt.Println(planetMarkII)
}