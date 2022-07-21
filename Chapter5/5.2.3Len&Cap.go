package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	s1 := arr[2:8]
	fmt.Println("len(s1)=", len(s1), ",cap(s1)=", cap(s1))
	s2 := arr[4:7]
	fmt.Println("len(s1)=", len(s2), ",cap(s1)=", cap(s2))
	s3 := arr[2:5]
	fmt.Println("len(s1)=", len(s3), ",cap(s1)=", cap(s3))
	s4 := arr[0:2]
	fmt.Println("len(s1)=", len(s4), ",cap(s1)=", cap(s4))
	/*
		len 计算的是当前切片里面有多少元素
		cap 计算的是从当前切片左端地址到数组右端地址的大小
			也就是说如果数组大小为11 哪怕切片后只有idx=0的元素
			该切片的大小也为11
			如果从下标3开始切片 则切片后的cap为8
	*/
}
