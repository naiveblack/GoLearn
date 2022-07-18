package main

import "fmt"

// 求1-100的和
func summation() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 求1-100之间3的倍数的和
func summation2() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

// 32米竹竿 每次截1.5米 几次后能小于4米
func cutBamboo() {
	cnt := 0
	for i := 32.0; i >= 4; i -= 1.5 { // i需要定义为float类型 i := 32报错
		cnt++
	}
	fmt.Println(cnt)
}
func main() {
	summation()
	summation2()
	cutBamboo()
}
