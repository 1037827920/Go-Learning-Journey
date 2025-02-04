// Description: 由结构体组成的slice

package main

import "fmt"

func main() {
	type location struct {
		name string
		lat float64
		long float64
	}	

	locations := []location {
		{"Bradbury Landing", 4.5895, 137.4417},
		{"Columbia Memorial Station", 14.5684, 175.472636},
		{"Challenger Memorial Station", -1.9462, 354.4734},
	}

	fmt.Println(locations)
}