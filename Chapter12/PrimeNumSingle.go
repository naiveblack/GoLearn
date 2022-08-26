package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 0
	start := time.Now().Unix()
	for i := 1; i <= 80; i++ {
		flag := true
		//for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			cnt += 1
			//fmt.Println("素数：", i)
		}
	}
	end := time.Now().Unix()
	fmt.Println("所耗时间为：", end-start)
	fmt.Println(cnt)
}
