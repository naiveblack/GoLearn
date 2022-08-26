package main

import (
	"fmt"
)

func f1() {
	for {
		fmt.Println("call f1...")
	}
}

func f2() {
	fmt.Println("call f2...")
}

/*
	注释掉f1之后发生死锁的问题

	没注释掉f1的时候，还有其他goroutine在跑，
	main goroutine中阻塞等待数据的<-ch有可能可以等到数据，所以没有死锁；
	但是注释掉f1之后，当f2执行完之后就只剩main goroutine在跑，
	而这个main goroutine又阻塞了，
	所以就永远不可能有其它goroutine可以发送数据到ch中，
	这个main goroutine就会永远阻塞了，就产生死锁了
*/
func main() {
	//go f1()
	//go f2()
	//ch := make(chan int)
	//<-ch
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("上当了")
	}()
	data, ok := <-ch
	fmt.Println(data, ok)

	// 死锁
	//ch := make(chan int)
	//ch <- 1
	//data, ok := <-ch
	//fmt.Println(data, ok)
	//
	//time.Sleep(2 * time.Second)
}
