// Description: 通过指针对数组的元素进行修改
// 隐式指针，比如map就是，map在被赋值或传递的时候不会被复制
// slice也是

package main

import "fmt"

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func main() {
	var board [8][8]rune
	reset(&board)

	fmt.Printf("%c\n", board[0][0])
}