package main

import (
	"fmt"
	"time"
)

/*
	select语句类似switch语句
	但select是从中随机执行一个可执行的case
	没有case可执行时将会阻塞
	每个case都必须是一个通信
	如果有default语句，则阻塞时执行该语句
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		data := <-ch1
		fmt.Println("ch1: ", data)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		data := <-ch2
		fmt.Println("ch2: ", data)
	}()

	select {
	case ch1 <- 100: //阻塞
		close(ch1)
		fmt.Println("ch1中写入数据...")
	case ch2 <- 200: //阻塞
		close(ch2)
		fmt.Println("ch2中写入数据...")
	case <-time.After(2 * time.Second):
		fmt.Println("执行延迟通道")
	default:
		fmt.Println("执行default")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("main over ...")
}
