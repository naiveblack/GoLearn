package main

import "fmt"

type Human2 struct {
	name, phone string
	age         int
}
type Student2 struct {
	Human2 //匿名字段
	school string
}
type Employee2 struct {
	Human2  //匿名字段
	company string
}

func main() {
	s1 := Student2{Human2{"Daniel", "15012345678", 13}, "十一中学"}
	e1 := Employee2{Human2{"Steven", "17812345678", 35}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
}
func (h *Human2) SayHi() {
	fmt.Printf("大家好!我是 %s ，我%d 岁，我的联系方式是：%s\n", h.name, h.age, h.phone)
}

//Student 的 method 重写 Human 的 method
func (s *Student2) SayHi() {
	fmt.Printf("大家好! 我是 %s ，我%d 岁，我在%s 上学，我的联系方式是：%s\n", s.name, s.age,
		s.school, s.phone)
}

//Employee 的 method 重写 Human 的 method
func (e *Employee2) SayHi() {
	fmt.Printf("大家好! 我是 %s ，我%d 岁，我在%s 工作，我的联系方式是：%s\n", e.name, e.age,
		e.company, e.phone)
}

/*
	重写就类似方法中接受者不同，可以取相同的方法名
*/
