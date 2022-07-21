package main

import "fmt"

// 操作指针改变变量的数值
func main() {
	b := 3214
	a := &b
	fmt.Println("b的地址为: ", a)
	fmt.Println("a的值为: ", *a)
	*a++
	fmt.Println("b的值改变后：", b)
}
