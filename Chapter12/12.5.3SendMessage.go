package main

func main() {
	// 创建空指针型通道
	ch := make(chan interface{})
	// 将0通过通道发送
	ch <- 0
	// 发送字符串
	ch <- "hello world"
}
