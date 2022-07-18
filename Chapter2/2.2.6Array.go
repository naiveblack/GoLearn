package main

import "fmt"

func printArr1(x [5]int) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func printArr2(x [5]int) {
	for key, value := range x {
		fmt.Printf("下标为%d, 值为%d\n", key, value)
	}
}

func main() {
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	printArr1(b)
	printArr2(b)
}
