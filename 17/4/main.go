// Description: 打印struct

package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	curiosity := location{-4.5895, 137.4417}
	// + 可以打印出struct的字段名
	fmt.Printf("%+v\n", curiosity)
}