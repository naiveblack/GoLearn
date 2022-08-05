package main

import "fmt"

func printAdd(a, b int) {
	fmt.Println("两个数为：", a, ",", b, "\n相加得到的和：", a+b)
}

func main() {
	a := 4
	b := 3
	defer printAdd(a, b)
	a = 23
	b = 12
	printAdd(a, b)
}
