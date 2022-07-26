package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "缓缓飘落的枫叶像思念~"
	fmt.Println(s)
	fmt.Println("字节长度：", len(s))
	fmt.Println("===============")

	// 获取字符串引用
	for i, ch := range s {
		fmt.Printf("%d:%c \n", i, ch)
	}
	fmt.Println("===============")

	// 遍历所有字节
	fmt.Println([]byte(s))
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%X \n", i, ch)
	}
	fmt.Println("===============")

	// rune相当于go里面的char类型
	count := 0
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c \n", i, ch)
		count++
	}
	fmt.Println("字符长度：", count)
	fmt.Println("字符长度：", utf8.RuneCountInString(s))
	fmt.Println("===============")
	/*
		Go 语言的内建函数 len()，可以用来获取切片、字符串、通道（channel）等的长度。
		其中中文等特殊字符一般占3个字节（不同于Java的2）

		如果希望按习惯上的字符个数来计算，
		就需要使用 Go 语言中 UTF-8 包提供的 RuneCountInString() 函数，
		统计 Uncode 字符数量。
	*/
}
