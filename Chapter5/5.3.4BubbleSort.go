package main

import "fmt"

func main() {
	arr := []int{5, 2, 7, 8, 1, 6, 9, 0, 4, 3}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	iCount := 0
	jCount := 0
	for i := 0; i < len(arr)-1; i++ {
		flag := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
			iCount++
		}
		jCount++
		if flag {
			break
		}
	}
	fmt.Println("i循环的次数：", iCount)
	fmt.Println("j循环的次数：", jCount)
}
