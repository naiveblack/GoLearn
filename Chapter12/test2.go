package main

import (
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	ch := make(chan int, 1)
	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		time.Sleep(1)
		fmt.Println("exit goroutine")
		//wg.Done()
	}()
	fmt.Println("wait goroutine")
	<-ch

	fmt.Println("all done")
	time.Sleep(1000000000)
	//wg.Wait()
}
