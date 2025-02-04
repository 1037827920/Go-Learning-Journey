// Description: 将struct编码为json

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exit_on_error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// type location struct {
	//	// json包要求字段必须能导出，所以用大写开头 
	// 	// go语言中的json包要求struc的字段必须是驼峰命名规范
	// 	// 但在实际使用中，我们可能需要snake_case蛇形命名规范，这时候
	// 	// 可以为字段注明标签，这样json在编码的时候就能够按照标签里的样式修改字段名
	// 	Lat float64
	// 	Long float64
	// }
	type location struct {
		// json包要求字段必须能导出，所以用大写开头
		Lat float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exit_on_error(err)

	fmt.Println(string(bytes))
}