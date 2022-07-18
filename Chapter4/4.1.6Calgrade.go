package main

import "fmt"

// 4.1.6 可变参数
func main() {
	// 1、传入n个参数
	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
	fmt.Printf("共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg)
	// fmt.Println("共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg)
	// 输出为 共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f 4 310.3 77.575

	// 2、传切片作为参数
	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
	sum, avg, count = GetScore(scores...) // 通过"..."将slice中的参数传递给对应函数
	fmt.Printf("共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg)
}

func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}
