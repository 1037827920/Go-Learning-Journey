// Description: 指针和接口

package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l* laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	shout(martian{})
	// 指针也满足talker这个接口
	shout(&martian{})

	// 但是方法如果使用的是指针接收者，情况就有所不同
	pew := laser(2)
	shout(&pew)
	// 下面这种写法是错误的
	// shout(pew)
}