package main

import (
	"fmt"
	"time"
)

//定义一个结构体，表示自定义错误的类型
type MyError struct {
	When time.Time
	What string
}

//自定义错误类型实现 error 接口的方法 ：Error() string
func (e *MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}

//定义一个返回 error 的函数。求矩形的面积
func getArea(width, length float64) (float64, error) {
	errorMsg := ""
	// fmt.Sprintf 可以赋值一个变量，fmt.Printf 只能打印
	if width < 0 && length < 0 {
		errorMsg = fmt.Sprintf("长度:%v ，宽度:%v ，均为负数", length, width)
	} else if length < 0 {
		errorMsg = fmt.Sprintf("长度:%v ，出现负数", length)
	} else if width < 0 {
		errorMsg = fmt.Sprintf("宽度:%v ，出现负数", width)
	}

	if errorMsg != "" {
		return 0, &MyError{time.Now(), errorMsg}
	} else {
		return width * length, nil
	}
}

func main() {
	res1, err := getArea(-4, -6)
	if err != nil {
		fmt.Printf(err.Error())
		//输出：2020-03-28 13:57:24.8765849 +0800 CST m=+0.004970401 : 长度:-6 ，宽度:-4 ，均为负数
	} else {
		fmt.Println("面积是:", res1)
	}
}
