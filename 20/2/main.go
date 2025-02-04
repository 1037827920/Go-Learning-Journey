// Description: 接口类型，通过type来定义一个接口类型

package main

import (
	"fmt"
	"strings"
)

// talker可以是任意实现了下面接口的类型
// 接口类型名通常以er结尾
type talker interface {
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

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))
}