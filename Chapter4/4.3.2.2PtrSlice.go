package main

import "fmt"

// 切片属于引用传递
func main() {
	a := []int{1, 2, 3, 4}
	b := [2]int{2, 3}
	fmt.Printf("%T, %T\n", a, b)
	fmt.Printf("1、变量 a 的内存地址是：%p ，值为：%v \n\n", &a, a) //[1,2,3,4]
	fmt.Printf("切片型变量 a 内存地址是：%p \n\n", a)            //可以获取到地址，类似：0xc420018080
	//传值
	changeSliceVal(a)
	fmt.Printf("changeSliceVal 函数调用后：变量 a 的内存地址是：%p，值为：%v \n\n", &a, a)
	//传引用
	changeSlicePtr(&a)
	fmt.Printf("changeSlicePtr 函数调用后：变量 a 的内存地址是：%p，值为：%v \n", &a, a)
}

func changeSliceVal(a []int) {
	fmt.Printf("--changeSliceVal 函数内：值参数 a 的内存地址是：%p，值为：%v \n", &a, a)
	//[1,2,3,4]
	fmt.Printf("--changeSlicePtr 函数内：值参数 a 的内存地址是：%p \n", a)
	a[0] = 99
}

func changeSlicePtr(a *[]int) {
	fmt.Printf("--changeSlicePtr 函数内：指针 a 的内存地址是：%p ，值为：%v \n", &a, a)
	//&[1,2,3,4]
	(*a)[1] = 250
}
