package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Compare(a, b string) int
	// 按字典序比较a和b字符串的大小
	// a > b 返回 1; a < b 返回 -1; a = b 返回 0
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1

	// fac EqualFold(s, t string) bool
	// 判断两个utf8字符串是否相等，忽略大小写
	fmt.Println(strings.EqualFold("Go", "go")) // true

	// func Repeat(s string, count int) string
	// 将字符串s重复count次
	fmt.Println("QwQ" + strings.Repeat("T T", 3))

	// func Replace(s, old, new string, n int) string
	// 用new替换s中old的部分 替换n次 n小于0 则全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "ooo", -1))

	// func Join(a []string, sep string) string
	// 将a中字符串以sep为分割做连接
	s := []string{"被伤透的心", "能不能够", "继续爱我"}
	fmt.Println(strings.Join(s, ","))
}
