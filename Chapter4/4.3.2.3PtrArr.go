package main

import "fmt"

// 数组属于值传递
func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Printf("1、变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a)
	//传值
	changeArrayVal(a)
	fmt.Printf("2、changeArrayVal 调用后：变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a) // 不改变数据
	//传引用
	changeArrayPtr(&a)
	fmt.Printf("3、changeArrayPtr 调用后：变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a)
}

func changeArrayVal(a [4]int) {
	fmt.Printf("--changeArrayVal 内：值参数 a 的内存地址是：%p ，值为：%v \n", &a, a)
	a[0] = 99
}

func changeArrayPtr(a *[4]int) {
	fmt.Printf("--changeArrayPtr 内：指针参数 a 的内存地址是：%p ，值为：%v \n", &a, a)
	(*a)[1] = 250
}
