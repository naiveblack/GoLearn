package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	condition := false
	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子 goroutine 已经锁定。。。")
		fmt.Println("子 goroutine 更改条件数值，并发送通知。。")
		condition = true //更改数值
		cond.Signal()    //发送通知：一个 goroutine
		fmt.Println("子 gorutine。。。继续。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("子 groutine 解锁。。")
		cond.L.Unlock()
	}()
	cond.L.Lock()
	fmt.Println("main..已经锁定。。。")
	if !condition {
		fmt.Println("main.。即将等待。。。")
		//wait()
		// 1.wait 尝试解锁，
		// 2.等待--->当前的 groutine 进入了阻塞状态，等待被唤醒：signal(),broadcast()
		// 3.一旦被唤醒后，又会锁定
		cond.Wait()
		fmt.Println("main.被唤醒。。")
	}
	fmt.Println("main。。。继续")
	fmt.Println("main..解锁。。。")
	cond.L.Unlock()
}
