package main

import (
	"fmt"
	"time"
)

func main() {
	// 1、非缓冲通道
	ch1 := make(chan int)
	fmt.Println("非缓冲通道：", len(ch1), cap(ch1))
	go func() {
		data := <-ch1
		fmt.Println("获取数据：", data)
	}()
	ch1 <- 100 //阻塞
	time.Sleep(1)
	fmt.Println("写入数据 ok \n -----------------------")

	// 2、缓冲通道
	ch2 := make(chan int, 5)
	fmt.Println("缓冲通道：", len(ch2), cap(ch2))
	go func() {
		for data := range ch2 {
			//time.Sleep(1)
			fmt.Println("获取数据：", data)
		}
	}()
	ch2 <- 1
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 2
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 3
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 4
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 5
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 6 //没有写入进去 但是为啥没死锁呢...
	fmt.Println(len(ch2), cap(ch2))
	close(ch2)
	fmt.Println("main...over...\n-------------------------")

	// 3、缓冲通道
	ch3 := make(chan string, 5)
	fmt.Printf("%T\n", ch3)
	go sendData2(ch3)

	for data := range ch3 {
		fmt.Println("\t读取数据：", data)
	}
	fmt.Println("读取完毕...")
}

func sendData2(ch3 chan string) {
	for i := 1; i <= 10; i++ {
		ch3 <- fmt.Sprintf("data%d", i)
		fmt.Println("写入数据：", i)
		fmt.Println(len(ch3), cap(ch3))
	}
	close(ch3)
}
