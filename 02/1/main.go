// Description: Print和Println函数

package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")

	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365/687, "years old.")
}
