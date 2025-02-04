// Description: slice也可以用于切分字符串

package main

import "fmt"
 
func main() {
	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune)

	neptune = "Poseidon"
	fmt.Println(tune)
}