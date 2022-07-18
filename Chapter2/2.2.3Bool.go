package main

import "fmt"

func main() {
	var v1 bool
	fmt.Println(v1) // 初始值为false

	v1 = true
	fmt.Println(v1)

	v1 = (1 == 2) // 可以给bool类型赋值逻辑判断表达式
	fmt.Println(v1)

	v2 := (2 == 2)
	fmt.Println(v2)

	var a, b int
	a = 2
	b = 2
	v3 := (a == b)
	fmt.Println(v3)
}
