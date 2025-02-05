// Description: recover函数
// defer的动作会在函数返回前执行，即使发生了panic，仍会继续执行

package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// 程序不会发生停止，因为recover函数
	panic("I forgot my towel")
}