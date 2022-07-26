package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	play()
}

func play() {
	n := 6
	target := generateRandNum(10, 100)
	fmt.Println("您有", n, "次机会, 请输入猜测的数字")

	cnt := 0
	for cnt < n {
		cnt++
		inputNum := 0
		fmt.Scanln(&inputNum)
		if inputNum == target {
			// 正确
			fmt.Println("猜测正确！ 一共猜了", cnt, "次")
			return
		} else if inputNum > target {
			fmt.Println("猜测数字大了")
		} else if inputNum < target {
			fmt.Println("猜测数字小了")
		} else {
			fmt.Println("未知错误")
		}
	}
	fmt.Println("正确答案为", target, ",再接再厉")
	return
}

func generateRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
