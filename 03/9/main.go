// Description: 练习题，随机生成一个数，判断这个数是大于、小于还是等于50

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const result_num = 50
	var random_num = rand.Intn(100) + 1

	if random_num < result_num {
		fmt.Println("random_num is less than result_num")
	} else if random_num > result_num {
		fmt.Println("random_num is greater than result_num")
	} else {
		fmt.Println("random_num is equal to result_num")
	}
}
