// Description: 将指针作为方法的接收者

package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func (p* person) birthday() {
	p.age++
}

func main() {
	terry := &person{name: "Terry", superpower: "fly", age: 21}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := person{name: "Nathan", superpower: "teleport", age: 21}
	nathan.birthday()
	// 等价于
	// (&nathan).birthday()
	fmt.Printf("%+v\n", nathan)
}