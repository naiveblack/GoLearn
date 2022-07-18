package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
	非具名函数不会受到defer的影响。所以返回值为5
*/
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

/*
	具名函数，返回值带有名字。
	这样在执行defer的时候相当于修改了返回值的值。
	所以返回值为6。
*/
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

/*
	go中的return语句并不是原子性操作，一般是分为两步:
	1、将返回值赋值给一个变量
	2、执行RET指令。
	defer就执行在1之后，2之前。

	defer后面的函数在入栈的时候保存的是入栈那一刻的值，
	后期对x修改，并不会影响栈内函数的值。
	但是如果函数入栈的时候存入的执行变量x的指针。
	那么后期x值改变的时候，输出结果也会改变。
*/
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
