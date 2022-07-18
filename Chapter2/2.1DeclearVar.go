package main

import "fmt"

func main() {
	// 标准变量声明
	var str string
	// 多变量声明方式
	var (
		a int
		b string
		c []float32
		d func() bool
		e struct {
			x int
			y string
		}
	)
	fmt.Println(str, a, b, c, d, e)

	//变量初始化
	var str1 string = "this is string1"
	fmt.Println(str1)
	var str2 = "this is string2"
	fmt.Println(str2)
	str3 := "this is string3"
	fmt.Println(str3)

	//匿名变量
	//i, _ := noname() 	// '_'不占用命名空间 不会分配内存
}
