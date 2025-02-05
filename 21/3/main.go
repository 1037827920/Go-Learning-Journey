// Description: 指向结构体的指针
// 与字符串和数值不一样，复合字面量的前面可以放置&

package main

import "fmt"

func main() {
	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"
	// 等价于
	// (*timmy).superpower = "flying"

	fmt.Printf("%+v\n", timmy)
}
