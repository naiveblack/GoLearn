package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%T\n", wg) // sync.WaitGroup
	fmt.Println(wg)        // {{} 0 0}
	wg.Add(3)
	rand.Seed(time.Now().UnixNano())

	go printNum1(&wg)
	go printNum2(&wg)
	go printNum3(&wg)

	wg.Wait() // main goroutine 进入阻塞状态，当计数器为0后解除
	fmt.Println("main解除阻塞, main over...")
}

func printNum1(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Println("子 goroutine1, i : ", i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()

}

func printNum2(wg *sync.WaitGroup) {
	defer wg.Done() // 延迟执行wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("\t 子 goroutine2, i : ", i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
}

func printNum3(wg *sync.WaitGroup) {
	wg.Done() // 直接执行wg.Done
	/*
		当后续需要执行的程序复杂度没有另外两个复杂的时候，
		该行为不会影响整体的运行。
		但是当此处的运算复杂度高于另外两个，
		由于提前结束了wg的等待，当另两个完成后，主程序中的Wait就满足了条件
		主程序便会继续执行下去，自然此处不会输出到100000(输出到11左右)
	*/
	for i := 1; i <= 100000; i++ {
		fmt.Println("\t\t 子 goroutine3, i : ", i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
}
