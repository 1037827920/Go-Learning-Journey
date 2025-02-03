// Description: range函数，遍历集合的值

package main

import "fmt"

func main() {
	message := "Hello, Wolrd!"

	for _, c := range message {
		fmt.Printf("%c\n", c)
	}
}
