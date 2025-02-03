// Description: switch语句

package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is cavern entrance here and a path to the east")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further into the cave")
	case "enter case", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'")
	default:
		fmt.Println("Nothing happens")
	}
}