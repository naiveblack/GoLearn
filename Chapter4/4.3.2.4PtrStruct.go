package main

import "fmt"

type Teacher struct {
	name    string
	age     int
	married bool
	sex     int8
}

// 结构体属于
func main() {
	a := Teacher{"Steven", 35, true, 1}
	fmt.Printf("1、变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a)
	//传值
	changeStructVal(a)
	fmt.Printf("2、changeArrayVal 调用后：变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a)
	//传引用
	changeStructPtr(&a)
	fmt.Printf("3、changeArrayPtr 调用后：变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a)
}
func changeStructVal(a Teacher) {
	fmt.Printf("--changeArrayVal 内：值参数 a 的内存地址是：%p ，值为：%v \n", &a, a) // &a内存地址与主函数中的地址不同
	a.name = "Josh"
	a.age = 29
	a.married = false
}
func changeStructPtr(a *Teacher) {
	fmt.Printf("--changeArrayPtr 内：指针参数 a 的内存地址是：%p ，值为：%v \n", &a, a)

	(*a).name = "Daniel"
	(*a).age = 20
	(*a).married = false
}
