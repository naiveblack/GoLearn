package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[:]
	fmt.Printf("arr类型%T，s类型%T\n", arr, s)
	fmt.Println(s)
	s = arr[1:2]
	fmt.Println(s) // 即 切片的范围是左闭右开
	s = arr[3:]
	fmt.Println(s)
	s = arr[:2]
	fmt.Println(s)
}
