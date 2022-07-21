package main

import "fmt"

func main() {
	// 1、改变一个元素
	a := 58
	fmt.Println("函数调用前a的值为: ", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%x \n", &a)
	var b *int = &a
	change(b)
	fmt.Println("函数调用之后的 a 的值：", a)
	change2(a)
	fmt.Println("调用非指针函数后a的值为: ", a)

	fmt.Println("==========================")
	// 2、两数交换
	c := 90
	swap0(a, c) //不会发生数值的交换
	fmt.Println("非指针交换后a和c的值为: ", a, c)
	swap(&a, &c)
	fmt.Println("指针交换后a和c的值为: ", a, c)

	fmt.Println("==========================")
	// 3、切片（数组元素）
	arr := []int{1, 2, 3}
	modify0(arr) //可以进行数值的修改
	fmt.Println("非指针传递后 arr[0]为：", arr[0])
	modify(&arr)
	fmt.Println("指针传递后 arr[0]为：", arr[0])
}

func change(val *int) {
	*val = 15
}

func change2(val int) {
	val = 33
}

func swap0(x, y int) (int, int) {
	return y, x
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func modify0(sls []int) {
	sls[0] = 100
}

func modify(sls *[]int) {
	(*sls)[0] = 200
}
