package main

import "fmt"

func main() {
	// make(类型， 长度， 容量)
	numbers := make([]int, 0, 20)
	fmt.Println(numbers, len(numbers))
	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	count := copy(numbers1, numbers) //copy返回的是拷贝的数量
	fmt.Println("拷贝的个数为", count)
	printSlices("numbers1:", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100
	printSlices("numbers:", numbers)
	printSlices("numbers1:", numbers1)
	/*
		copy会复制切片元素，同时在新的地址创建一个切片元素
		即 新的切片元素和原始切片元素之间没有联系 数据不会同时更改
	*/

	// 利用append的方法 删除其中的一个元素
	numbers = append(numbers[:2], numbers[3:]...)
	printSlices("numbers:", numbers)
}

func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x), cap(x), x)
}
