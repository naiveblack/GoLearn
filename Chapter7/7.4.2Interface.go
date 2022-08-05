package main

import "fmt"

type Phone interface {
	call()
}
type AndroidPhone struct {
}
type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，我可以打电话!")
}
func (i IPhone) call() {
	fmt.Println("我是苹果手机，我可以打电话!")
}

/*
	只要类型T实现了接口I中的所有方法，则表明其自动实现了接口I
	如果func (i IPhone) call() 没有实现
	则 phone = new(IPhone) 报错
*/
func main() {
	//定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
}
