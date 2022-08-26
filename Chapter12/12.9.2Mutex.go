package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var tickets = 100
var mutex sync.Mutex

func main() {
	/*
		互斥锁模拟售票
	*/
	var wg sync.WaitGroup
	wg.Add(4)

	go saleTickets("售票口A", &wg)
	go saleTickets("售票口B", &wg)
	go saleTickets("售票口C", &wg)
	go saleTickets("售票口D", &wg)

	wg.Wait()
	fmt.Println("车票已售空")
}

func saleTickets(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name, ":", tickets)
			tickets -= 1
		} else {
			fmt.Println(name, "结束售票")
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
}
