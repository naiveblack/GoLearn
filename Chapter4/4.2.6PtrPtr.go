package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	a = 1234
	// ptr指向元素a的地址
	ptr = &a
	fmt.Println("ptr", ptr)
	// pptr 指向指针ptr的地址
	pptr = &ptr
	fmt.Println("pptr", pptr)
	// 获取a，ptr，pptr里面的值
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
