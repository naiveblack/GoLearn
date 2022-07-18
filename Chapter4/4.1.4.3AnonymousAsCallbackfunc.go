package main

import "fmt"
import "math"

// 4.1.4 3 匿名函数作为回调函数
func main() {
	// 调用函数， 对每个元素进行求平方操作
	arr := []float64{1, 9, 16, 25, 30}
	visit(arr, func(v float64) { // 调用visit方法后，利用此处的func对数组中的元素进行处理
		v = math.Sqrt(v)
		fmt.Printf("%.2f\n", v)
	})

	// 调用函数，对每个元素进行求平方操作
	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.2f\n", v)
	})
}

// 遍历切片元素
func visit(list []float64, f func(float64)) {
	for _, value := range list {
		f(value)
	}
}
