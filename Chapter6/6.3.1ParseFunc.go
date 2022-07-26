package main

import (
	"fmt"
	"strconv"
)

func main() {
	// func Atoi(s string) (int, error
	// 返回转int类型的结果 是ParseInt(s, 10, 0)的缩写
	v := "20"
	fmt.Println(strconv.Atoi(v))
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	/*
		s表示待输入的字符串
		base表示进制 认定字符串为多少进制
			--如 s=111 base=2 输出7；base=3 输出13
		bitSize表示int的大小int int8 int 16 int32 int64
		用于限制int的范围 当转换后的十进制数超出范围 则报错
	*/
	v32 := "-522"
	//v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		//fmt.Println(err.Error())
		fmt.Printf("%T, %v\n", s, s)
	}

	// func ParseUint(s string, base int, bitSize int) (uint64, error)
	// 同理ParseInt 只不过是无符号
	//v32 := "-354634382"
	v3 := "522"
	if s, err := strconv.ParseUint(v3, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v3, 16, 32); err == nil {
		//fmt.Println(err.Error())
		fmt.Printf("%T, %v\n", s, s)
	}

	// func ParseFloat(s string, bitSize int) (float64, error)
	/*
		bitSize有32和64
		当bitSize=32时，结果仍具有float64类型
		但可以在不更改其值的情况下转为float32
	*/
	var vf = "3.1415926535"
	if s, err := strconv.ParseFloat(vf, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat(vf, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	// func ParseBool(str string) (bool, error)
	// 返回字符串表示的布尔值 限定输入 其余都会报错
	// 包括：0, f, F, FALSE, false, False
	// 		1, t, T, TRUE , true , True
	vb := "True"
	if s, err := strconv.ParseBool(vb); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
