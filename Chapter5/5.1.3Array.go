package main

import "fmt"

func main() {
	var a = [5][2]int{{1, 0}, {2, 4}, {6, 2}, {3, 7}, {0, 5}}
	fmt.Println(len(a))    // 长度是数组有多少行
	fmt.Println(len(a[0])) // 某一行有多少元素
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("%d \t", a[i][j])
		}
		println()
	}
}
