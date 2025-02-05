// Description: 使用select处理多个channel
// 注意：即使已经停止等待goroutine，但只要main函数还没返回，仍在运行的goroutine将会继续
// 占用内存

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	// 超时时间是2s
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		// 如果有数据从channel中读取，就执行case
		case gopherId := <-c:
			fmt.Println("gopher ", gopherId, " has finished sleeping")
		// 如果超时，就执行下面的case
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}

}

func sleepyGopher(id int, c chan int) {
	// 等待随机时间0-4秒
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	// 向channel发送数据
	c <- id
}
