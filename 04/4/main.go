// Description: 生成随机年份，能判断是否是闰年，使用for循环生成和展示10个日期

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for i := 0; i < 10; i++ {

		year := rand.Intn(3000) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31

		// 闰年判断
		is_leap_year := year%400 == 0 || (year%4 == 0 && year%100 != 0)

		switch month {
		case 2:
			if is_leap_year {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(era, year, month, day)
	}

}
