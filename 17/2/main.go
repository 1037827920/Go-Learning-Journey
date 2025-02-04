// Description: 通过了type复用结构体

package main

import "fmt"

type location struct {
	lat float64 // 纬度
	lon float64 // 经度
}

func main() {


	var spirit location
	spirit.lat = -14.5684
	spirit.lon = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.lon = 354.4734

	fmt.Println(spirit, opportunity)
}