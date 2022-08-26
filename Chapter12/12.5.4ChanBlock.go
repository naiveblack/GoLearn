package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 chan int
	fmt.Println(ch1)
	fmt.Printf("%T\n", ch1)

	ch1 = make(chan int)
	fmt.Println(ch1)
	ch2 := make(chan bool)

	go func() {
		fmt.Println("子goroutine...")
		data, ok := <-ch1 // 阻塞，从通道中读取数据
		time.Sleep(1 * time.Second)
		fmt.Println("子goroutine从通道中读取到main传来的数据是：", ok, data)
		ch2 <- true
	}()

	ch1 <- 100 // 阻塞式， 向通道中写入数据
	<-ch2      // 防止main goroutine执行完毕后退出，如果main退出则不执行别的goroutine
	fmt.Println("main over")
}
