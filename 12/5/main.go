// Description: 闭包，由于匿名函数封闭并包围作用域中的变量而得名

package main

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	// 虽然calibrate已经返回，但实际执行sensor()时，仍然可以访问两个传入的参数
	sensor := calibrate(realSensor, 5)
	println(sensor())
}