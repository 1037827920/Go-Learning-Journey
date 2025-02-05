// Description: channel

package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建channel
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	// 主线路
	for i := 0; i < 5; i++ {
		// 从channel接收数据
		gopherId := <-c
		fmt.Println("gopher ", gopherId, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	// 向channel发送数据
	c <- id
}
