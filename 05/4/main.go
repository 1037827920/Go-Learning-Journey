// Description: 随机地将0.05美元，0.10美元和0.25美元放入一个空的储蓄罐
// 直到里面至少有20美元。每次村矿显示存钱罐的金额，并以适当的宽度和精度格式化

package main

import (
	"fmt"
	"math/rand"
)


func main() {
	coins := []float64{0.05, 0.10, 0.25}
	var sum float64

	for sum < 20.0 {
		coin := coins[rand.Intn(len(coins))]
		sum += coin
		fmt.Printf("coin: %5.2f, sum: %5.2f\n", coin, sum)
	}
}