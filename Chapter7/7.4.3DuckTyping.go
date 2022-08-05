package main

import "fmt"

type ISayHello interface {
	SayHello()
}
type Persond struct{}
type Duck struct{}
type Duck2 struct{}

func (person Persond) SayHello() {
	fmt.Printf("Hello!")
}
func (duck Duck) SayHello() {
	fmt.Printf("ga ga ga!")
}
func greeting(i ISayHello) {
	i.SayHello()
}
func main() {
	//person := Person{}
	//duck := Duck{}
	person := new(Persond)
	duck := new(Duck)
	//以下输出跟接口没有关系
	fmt.Println("非接口调用形式")
	person.SayHello()
	duck.SayHello()
	fmt.Println("\n---------------------")
	//定义接口变量。
	fmt.Println("接口调用形式")
	var i ISayHello
	i = person
	greeting(i)
	i = duck
	greeting(i)
	//可否将一个未实现接口方法的结构体对象赋值给接口呢？ --不可以
	//i = new(Duck2)
}
