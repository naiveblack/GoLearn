package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- 100
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch2 <- 200
		}
	}()
	//time.Sleep(3 * time.Second) // 输出ch1
	for {
		select {

		case data := <-ch1:
			fmt.Println("ch1中数据：", data)
		case data := <-ch2:
			fmt.Println("ch2中数据：", data)
			//default:
			//	fmt.Println("执行了default")
		}
	}

}
