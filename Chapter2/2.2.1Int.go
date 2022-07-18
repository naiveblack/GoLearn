package main

import "fmt"

// 数据类型
func main() {
	var v int8
	v1 := 10
	//v = v1  // 报错 int与int8类型不能自动进行转换
	v = int8(v1) //强制类型转换
	fmt.Println(v)

	var v2 byte // uint8和byte是同一类型
	var v3 uint8
	v3 = 10
	v2 = v3
	fmt.Println(v2)
}
