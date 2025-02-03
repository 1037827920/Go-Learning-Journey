// Description: 使用big package

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 从int64类型的值创建Big Int
	light_speed := big.NewInt(299792)
	seconds_per_day := big.NewInt(86400)
	fmt.Println(light_speed, seconds_per_day)

	// 但是如果本身的值超过了int64的范围，就不能上面那样写了
	distance := new(big.Int)
	// 第一个参数是字符串（即值），第二个参数是进制
	distance.SetString("24000000000000000000000", 10)
	fmt.Println(distance)
}