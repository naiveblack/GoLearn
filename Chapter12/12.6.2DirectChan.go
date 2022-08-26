package main

import (
	"fmt"
	"time"
)

func main() {
	// 1、双向通道
	ch1 := make(chan string)
	go fun1(ch1)

	data := <-ch1
	fmt.Println("main, 接受到的数据：", data)

	ch1 <- "go go go"
	ch1 <- "blockchain"

	go fun2(ch1)
	go fun3(ch1)
	time.Sleep(1 * time.Second)
	fmt.Println("main over...")
}

func fun1(ch1 chan string) {
	ch1 <- "this is function 1"
	data := <-ch1
	data2 := <-ch1
	fmt.Println("function 1 ：", data, data2)
}

func fun2(ch1 chan<- string) {
	// 只能写入的通道
	ch1 <- "just can write"
}

func fun3(ch1 <-chan string) {
	// 只能读取的通道
	data := <-ch1
	fmt.Println("just read: ", data)
}
