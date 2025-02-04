// Description: 使用make函数对map进行预分配

package main

import "fmt"

func main() {
	temperature := make(map[float64]int, 8)

	fmt.Println(temperature)

	fmt.Printf("len=%v", len(temperature))
}