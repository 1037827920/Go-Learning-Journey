// Description: 随机地将5美分，10美分和25美分放入一个空的储蓄罐
// 直到里面至少有20美元。每次村矿显示存钱罐的金额，并以适当的宽度和精度格式化
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	coins := []int{5, 10, 25}
	var sum float64

	for sum < 20.0 {
		coin := coins[rand.Intn(len(coins))]
		sum += float64(coin) / 100
		fmt.Printf("coin: %5.2f, sum: %5.2f\n", float64(coin)/100, sum)
	}
}