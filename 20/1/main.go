// Description: 接口

package main

import (
	"fmt"
	"strings"
)

// t可以是任意实现了下面接口的类型
var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}