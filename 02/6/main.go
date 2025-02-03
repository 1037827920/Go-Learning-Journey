// Description: 自增变量

package main

import "fmt"

func main() {
	var age = 41
	age = age + 1
	age += 1
	age++ // go没有++age
	fmt.Println(age)
}