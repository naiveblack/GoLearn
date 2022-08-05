package main

import (
	"fmt"
)

type Dog struct {
	name  string
	color string
	age   int8
	kind  string //品种
}

func main() {
	//实现结构体的深拷贝
	//struct 的数据类型：值类型，所以默认的复制就是深拷贝
	d1 := Dog{"豆豆", "黑色", 2, "二哈"} //Dog
	fmt.Printf("d1：%T , %v , %p \n", d1, d1, &d1)
	d2 := d1 //深拷贝 dog
	fmt.Printf("d2：%T , %v , %p \n", d2, d2, &d2)
	//修改 d2，d1 是否也发生变化？
	d2.name = "毛毛"
	fmt.Println("d2 修改后=", d2)
	fmt.Println("d1=", d1)
	fmt.Println("------------------------")

	//实现结构体的浅拷贝：直接拷贝指针地址实现浅拷贝
	d3 := &d1
	fmt.Printf("d3：%T , %v , %p \n", d3, d3, d3)
	d3.kind = "萨摩耶"
	d3.color = "白色"
	d3.name = "球球"
	fmt.Println("d3 修改后=", d3)
	fmt.Println("d1=", d1)
	fmt.Println("------------------------")

	//实现结构体的浅拷贝
	//拷贝通过 new 函数实例化的对象 new函数返回的是一个指针
	d4 := new(Dog) //*Dog
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 1
	d4.kind = "巴哥犬"
	d5 := d4 //*Dog
	fmt.Printf("d4：%T , %v , %p \n", d4, d4, d4)
	fmt.Printf("d5：%T , %v , %p \n", d5, d5, d5)
	//修改 d2，d1 是否也发生变化？
	d5.color = "金色"
	d5.kind = "金毛"
	fmt.Println("d5 修改后=", d5)
	fmt.Println("d4=", d4)
	fmt.Println("------------------------")
}
