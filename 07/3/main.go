// Description: 一直大矮星距离地球为23600000000000000公里，光速为300000公里/秒。将此距离转换为光年

package main

import "fmt"

func main() {
	fmt.Println(23600000000000000 / 300000 / 60 / 60 / 24 / 365)
}