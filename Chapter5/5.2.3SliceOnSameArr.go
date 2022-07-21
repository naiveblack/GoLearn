package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	nums1 := arr[:]
	nums2 := arr[:]
	nums3 := arr[1:]
	fmt.Printf("%p, %p, %p\n", nums1, nums2, nums3)
	nums1[2] = 90
	nums3[0] = 77
	arr[3] = 129
	fmt.Println(nums1, nums2, nums3)
	fmt.Println(arr)
	/*
		通过arr切片操作得到的切片元素的地址就是数组的地址
		对数组或者对任意切片元素进行更改 都会影响到数组和切片
	*/
}
