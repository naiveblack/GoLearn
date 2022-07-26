package main

import "fmt"

func main() {
	username := ""
	age := 0
	fmt.Scanln(&username, &age)
	fmt.Println("账号信息为：", username, age)
	// %q 带双引号的字符串或单引号的rune
	fmt.Printf("用户名是：%q ， 年龄是：%d \n", username, age)
	fmt.Printf("用户名是：%s ， 年龄是：%d \n", username, age)
	fmt.Println(&username)
}
