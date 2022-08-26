package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func getChars(s string) {
	for _, v := range s {
		fmt.Printf("%c at time %v\n", v, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getNumbers(s []int) {
	for _, v := range s {
		fmt.Printf("%d at time %v\n", v, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	fmt.Println("main execution started at time:", time.Since(start))
	go getChars("极光掠过天边")
	go getNumbers([]int{1, 2, 3, 4, 5, 6, 7})
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\nmain execution stopped at time:", time.Since(start))
}
