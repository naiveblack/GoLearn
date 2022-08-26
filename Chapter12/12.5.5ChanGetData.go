package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go sendData(ch1)
	// 循环接收数据 方式1
	for {
		data := <-ch1
		//通道关闭时，通道输出的为数据类型的默认值。string为""
		if data == "" {
			break
		}
		fmt.Println("从通道中读取数据 方式1：", data)
	}
}

func sendData(ch1 chan string) {
	for i := 1; i <= 10; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}
	fmt.Println("发送数据完毕...")
	close(ch1)
}
