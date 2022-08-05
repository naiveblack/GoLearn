package main

import "fmt"

func main() {
	s1 := []int{53, 68, 43, 75, 22}
	largest(s1)
}

func finished() {
	fmt.Println("结束")
}

func largest(s []int) {
	defer finished()
	fmt.Println("寻找最大值：")
	maxx := s[0]
	for _, v := range s {
		if v > maxx {
			maxx = v
		}
	}
	fmt.Println("最大值为：", maxx)
}
