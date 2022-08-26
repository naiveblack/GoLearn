package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d，尝试读锁定。。\n", i)
			rwm.RLock()
			fmt.Printf("goroutine %d，已经读锁定了。。\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d,读解锁。。\n", i)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main..尝试写锁定。。")
	rwm.Lock()
	fmt.Println("main。。已经写锁定了。。")
	rwm.Unlock()
	fmt.Printf("main。。写解锁。。。")
}
