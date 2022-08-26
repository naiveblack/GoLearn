package main

import (
	"fmt"
	"time"
)

func printS() {
	time.Sleep(100)
	fmt.Println("我会发着呆，然后微微笑，接着紧紧闭上眼")
}

func main() {
	fmt.Println("main execution started")
	go printS()
	time.Sleep(150)
	fmt.Println("main execution stopped")
}
