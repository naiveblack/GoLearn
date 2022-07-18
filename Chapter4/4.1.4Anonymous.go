package main

import "fmt"

func main() {
	// 在定义时 调用匿名函数
	func(data int) {
		data = data * 10
		fmt.Println("hello", data)
	}(100)

	// 将匿名函数赋值给变量
	f := func(data string) {
		fmt.Println(data)
	}
	f("hello, world")
}
