package main

import "fmt"

// 判断一年中的某个月有多少天
func checkDays() {
	year := 2008
	month := 2
	days := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d年%d月的天数为%d\n", year, month, days)
}

// 判断某个接口变量实际存储的变量类型
func checkType() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是:%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

func main() {
	checkDays()
	checkType()
}
