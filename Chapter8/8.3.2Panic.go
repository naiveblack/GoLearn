package main

import "fmt"

func main() {
	funA()
	funB()
	funC()
	fmt.Println("main...over....")
}
func funA() {
	fmt.Println("我是函数 funA()...")
}
func funB() { //外围函数
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "恢复啦。。。")
		}
	}()
	fmt.Println("我是函数 funB()...")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB 函数，恐慌啦。。。") //打断程序的执行。。
		}
	}
	//当外围函数中的代码引发运行恐慌时，
	//只有其中所有的延迟函数都执行完毕后，
	//该运行时恐慌才会真正被扩展至调用函数。
	//即恐慌会中断程序，但是defer会继续执行。
}
func funC() {
	defer func() {
		fmt.Println("func 的延迟函数。。。")
		//if msg := recover(); msg != nil {
		// fmt.Println(msg, "恢复啦。。。")
		//}
		fmt.Println("recover 执行了", recover())
	}()
	fmt.Println("我是函数 funC()。。")
	panic("funC 恐慌啦。。")
}
