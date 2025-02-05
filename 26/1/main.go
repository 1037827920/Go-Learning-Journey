// Description: 以下程序会发生死锁

package main

func main() {
	c := make(chan int)
	<-c
}