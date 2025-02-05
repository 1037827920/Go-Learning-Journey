// Description: slice在没有使用复合字面值或make预分配的情况下，默认值是nil
// map、接口也同理

package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s == nil)
}