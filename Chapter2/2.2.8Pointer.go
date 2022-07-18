package main

import "fmt"

func main() {
	a := 10
	var p *int
	p = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	//定义一个指向指针的指针
	var ptr **int
	ptr = &p
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(**ptr)
}
