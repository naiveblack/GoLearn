package main

import "fmt"

// 4.1.7 递归函数
func main() {
	// 1、通过递归实现阶乘的计算
	fmt.Println(fac(5))
	fmt.Printf("%d\n", fac(5))
	// fmt.Printf(fac(5)) // 报错

	// 2、通过循环实现阶乘的计算
	fmt.Println(getFac(5))
}

func fac(n int) int {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}

func getFac(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}
