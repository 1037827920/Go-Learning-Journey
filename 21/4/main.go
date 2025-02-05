// Description: 指向数组的指针
// Slice和map前面也可以放置&操作符，但是go并没有为它们实现自动解引用的功能

package main

import "fmt"

func main() {
	superpowers := &[3]string{"flight", "invisibility", "super strength"}

	// 也不需要进行解引用即可直接使用
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
}