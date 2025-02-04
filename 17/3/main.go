// Description: 通过复合字面值初始化struct

package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	// 第一种，可以忽略字段定义的顺序进行初始化，甚至可以不初始化一些字段，采用默认值
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	// 第二种，按照字段定义的顺序进行初始化
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
}