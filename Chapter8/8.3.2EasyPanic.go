package main

import "fmt"

func main() {
	f()

	fmt.Println("main end")
}

/*
	捕获代码recover逻辑需要放在defer语句中，
	否则函数将返回nil，看不到任何效果，
	因为defer语句在panic后也会被执行到。
	recover里面获取到的就是panic里面的字符串。
*/
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()

	ff()

	fmt.Println("f end")
}

func ff() {
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")
}
