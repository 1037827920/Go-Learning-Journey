// Description: struct的复制，是复制值，即多一份副本

package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.5859, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.0106

	fmt.Println(bradbury, curiosity)
}