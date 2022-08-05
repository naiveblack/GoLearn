package main

import (
	"fmt"
	"math"
	"strings"
)

//1.定义一个接口
type Shape interface {
	perimeter() float64
	area() float64
}

//2.矩形
type Rectangle struct {
	a, b float64
}

//3.三角形
type Triangle struct {
	a, b, c float64
}

//4.圆形
type Circle struct {
	radius float64 //半径
}

//实现接口的方法
func (r Rectangle) perimeter() float64 {
	return 2 * (r.a + r.b)
}
func (r Rectangle) area() float64 {
	return r.a * r.b
}
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.perimeter() / 2 //半周长
	//海伦公式
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

//测试函数
func testShape(s Shape) {
	fmt.Printf("周长：%.2f ， 面积：%.2f \n", s.perimeter(), s.area())
}
func main() {
	var s Shape
	s = Rectangle{3, 4}
	testShape(s)
	getType(s)
	getType2(s)
	fmt.Println(strings.Repeat("=", 15))
	s = Triangle{3, 4, 5}
	testShape(s)
	getType(s)
	getType2(s)
	fmt.Println(strings.Repeat("=", 15))
	s = Circle{1}
	testShape(s)
	getType(s)
	getType2(s)
}

//接口对象转型——方式 1 利用instance获取类型
func getType(s Shape) {
	if instance, ok := s.(Rectangle); ok {
		fmt.Printf("矩形：长度为%.2f ， 宽为%.2f ，\t", instance.a, instance.b)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形：三边分别为%.2f ，%.2f ， %.2f ，\t", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形：半径为%.2f ，\t", instance.radius)
	}
	fmt.Println()
}

//接口对象转型——方式 2 遍历接口对象的类型
func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle:
		fmt.Printf("矩形：长度为%.2f ， 宽为%.2f ，\t", instance.a, instance.b)
	case Triangle:
		fmt.Printf("三角形：三边分别为%.2f ，%.2f ， %.2f ，\t", instance.a, instance.b, instance.c)
	case Circle:
		fmt.Printf("圆形：半径为%.2f ，\t", instance.radius)
	}
	fmt.Println()
}
