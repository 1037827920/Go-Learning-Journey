// Description: goroutine

package main

import (
	"fmt"
	"time"
)

func main() {
	// 这个sleepyGopher()是并发执行的
	go sleepyGopher() // 分支线路
	// 如果没有睡这4s，主线路会先结束，导致分支线路也结束
	time.Sleep(4 * time.Second) // 主线路
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}