package main

import "fmt"
import "math"

func main() {
	// float32 和 float64的最大值 float64的精度大 误差相对较小
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	var f float32 = 1 << 24
	fmt.Println(f == f+1) //丢失精度
	fmt.Println(f)
	fmt.Println(f + 1)

	var f2 float32 = 123453.1415926
	fmt.Printf("%8.2f\n", f2)
	fmt.Printf("%3.1f\n", f2)
	fmt.Printf("%3.4f\n", f2) //进位

	test := 1.0
	test2 := 1 //不加小数点会默认为int类型
	fmt.Printf("%T, %T", test, test2)
}
