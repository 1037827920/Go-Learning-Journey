// Description: 定义struct

package main

import "fmt"

func main() {
	var curiosity struct {
		lat float64 // 经度
		long float64 // 纬度
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}