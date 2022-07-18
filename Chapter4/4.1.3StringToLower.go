package main

import (
	"fmt"
	"strings"
)

func main() {
	result := StringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	result2 := StringToLower2("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result2)
}

// 奇偶数依次显示大小写
func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func StringToLower(str string, f func(string) string) string {
	fmt.Printf("%T\n", f)
	return f(str)
}

type caseFunc func(string) string

func StringToLower2(str string, f caseFunc) string {
	fmt.Printf("%T\n", f)
	return f(str)
}
