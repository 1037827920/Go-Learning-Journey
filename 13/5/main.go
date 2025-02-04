// Description: 二维数组

package main

import "fmt"

func main() {
	var board [8][8]string

	board[0][0] = "r"
	board[0][1] = "n"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Println(board)
}