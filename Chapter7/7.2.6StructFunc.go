package main

import (
	"fmt"
)

type Flower struct {
	name  string
	color string
}

func main() {
	//测试结构体作为参数
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1：%T , %v , %p \n", f1, f1, &f1)
	fmt.Println("----------------------")

	//将结构体对象作为参数传递
	changeInfo1(f1)
	fmt.Printf("f1：%T , %v , %p \n", f1, f1, &f1)
	fmt.Println("----------------------")

	//将结构体指针作为参数传递
	changeInfo2(&f1)
	fmt.Printf("f1：%T , %v , %p \n", f1, f1, &f1)
	fmt.Println("----------------------")

	//结构体对象作为返回值
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println(f2, f3)
	f2.name = "杏花"
	fmt.Println(f2, f3)
	fmt.Printf("f2：%T , %v , %p \n", f2, f2, &f2)
	fmt.Printf("f3：%T , %v , %p \n", f3, f3, &f3)
	fmt.Println("----------------------")

	//结构体指针作为返回值
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println(f4, f5)
	f4.name = "桃花"
	fmt.Println(f4, f5)
	fmt.Printf("f4：%T , %v , %p \n", f4, f4, f4)
	fmt.Printf("f5：%T , %v , %p \n", f5, f5, f5)
	fmt.Println("----------------------")
	/*
		只有将结构体指针作为参数传递的时候会影响原来元素里的数据
		结构体或者结构体指针作为返回值的时候接收到的值地址都不同
	*/
}

//传结构体对象
func changeInfo1(f Flower) {
	f.name = "月季"
	f.color = "粉"
	fmt.Printf("函数内 f 修改后=%T , %v , %p \n", f, f, &f)
}

//传对象指针
func changeInfo2(f *Flower) {
	f.name = "蔷薇"
	f.color = "紫"
	fmt.Printf("函数内 f 修改后=%T , %v , %p \n", f, f, f)
}

//返回结构体对象
func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白"}
	return
}

//返回结构体指针
func getFlower2() (f *Flower) {
	f = &Flower{"芙蓉", "红"}
	return
}
