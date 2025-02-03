// Description: 声明函数

package main

import "fmt"

// 如果首字母大写说明是可以被其他包访问的
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	// 这里传入的kelvin是一个副本，所以不会影响原来的值
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°K is ", celsius, "°C")

	kelvin = 0
	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "°K is ", fahrenheit, "°F")
}