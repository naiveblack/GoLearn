package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	name    string
	age     int
	married bool
	sex     int8
}

func main() {
	var s1 = Student{"Steven", 35, true, 1}
	var s2 = Student{"Sunny", 20, false, 0}
	var a *Student = &s1 // a b 存储的为s1 s2的地址
	var b *Student = &s2
	fmt.Println("==========================")
	fmt.Printf("s1类型为%T，值为%v\n", s1, s1)
	fmt.Printf("s2类型为%T，值为%v\n", s2, s2)
	fmt.Println("==========================")
	fmt.Printf("a类型为%T，值为%v\n", a, a)
	fmt.Printf("b类型为%T，值为%v\n", b, b)
	fmt.Println("==========================")
	fmt.Println(s1.name, s1.age, s1.married, s1.sex)
	fmt.Println(a.name, a.age, a.married, a.sex)
	fmt.Println(s2.name, s2.age, s2.married, s2.sex)
	fmt.Println(b.name, b.age, b.married, b.sex)
	fmt.Println("==========================")
	fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
	fmt.Println((*b).name, (*b).age, (*b).married, (*b).sex)
	fmt.Println("==========================")
	// &取地址符号 获取指针a b的地址
	fmt.Printf("&a 类型为%T，值为%v\n", &a, &a)
	fmt.Printf("&b 类型为%T，值为%v\n", &b, &b)
	fmt.Println(&s1, a, &a)
	fmt.Println("==========================")
	fmt.Println(&a.name, &a.age, &a.married, &a.sex)
	fmt.Println(&b.name, &b.age, &b.married, &b.sex)
	fmt.Println(&(s1.name))
	fmt.Println(&(a.name)) // a->name
	fmt.Println("==========================")
	ptr := unsafe.Pointer(a) //将a转换成Pointer类型
	fmt.Println(ptr)
	targetAddress := uintptr(ptr) + 32 //得到了b的地址
	fmt.Println(targetAddress)
	newPtr := unsafe.Pointer(targetAddress) //将新地址转换成Pointer类型
	fmt.Println(newPtr)
	value := (*Student)(newPtr)
	fmt.Println(value)
}
