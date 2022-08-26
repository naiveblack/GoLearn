package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	//写入数据：生产者
	go func() {
		for i := 1; i <= 20; i++ {
			ch1 <- i
			fmt.Println("写入数据：", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch1)
	}()
	//读取数据：消费者
	go func() {
		for data := range ch1 {
			fmt.Println("\t1 号消费者：", data)
			// time.Millisecond 显示日期表示的毫秒部分
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		ch2 <- true
	}()
	go func() {
		for data := range ch1 { //1
			fmt.Println("\t2 号消费者：", data)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		ch2 <- true
	}()
	<-ch2
	fmt.Println("main...over...")
}
