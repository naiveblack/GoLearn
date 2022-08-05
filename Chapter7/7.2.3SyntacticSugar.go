package main

import (
	"fmt"
	"strings"
)

type Emp struct {
	name string
	age  int8
	sex  byte
}

func main() {
	emp1 := new(Emp)
	fmt.Printf("emp1: %T, %v, %p\n", emp1, emp1, emp1)
	(*emp1).name = "David"
	(*emp1).age = 30
	(*emp1).sex = 1
	// 语法简写形式，语法糖
	emp1.name = "Steven"
	emp1.age = 35
	emp1.sex = 1
	fmt.Println(emp1)
	fmt.Println(strings.Repeat("=", 15))
	SyntacticSugar()
}

func SyntacticSugar() {
	arr := [4]int{1, 2, 3, 4}
	arr2 := &arr
	fmt.Println((*arr2)[3])
	fmt.Println(arr2[3])

	// 切片中没有语法糖
	arr3 := []int{10, 20, 30, 40}
	arr4 := &arr3
	fmt.Println((*arr4)[3])
	//fmt.Println(arr4[3])
}
