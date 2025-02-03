//Description: 打印浮点数

package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	// 4是宽度，表示总共占4个字符，2是精度，表示小数点后保留2位
	fmt.Printf("%4.2f\n", third)
}