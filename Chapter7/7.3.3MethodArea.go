package main

import (
	"fmt"
	"math"
)

type Rectanglec struct {
	width, height float64
}
type Circlec struct {
	radius float64
}

func main() {
	r1 := Rectanglec{10, 4}
	r2 := Rectanglec{12, 5}
	c1 := Circlec{1}
	c2 := Circlec{10}
	fmt.Println("r1 的面积：", r1.area())
	fmt.Println("r2 的面积：", r2.area())
	fmt.Println("c1 的面积：", c1.area())
	fmt.Println("c2 的面积：", c2.area())
}

/*
	方法名相同 但是接受者不一样 则方法就不一样
	类似于类的行为的方法，传递不同的元素则是不同的方法。
*/

//该 method 属于 Rectangle 类型的对象
func (r Rectanglec) area() float64 {
	return r.width * r.height
}

//该 method 属于 Circle 类型的对象
func (c Circlec) area() float64 {
	return c.radius * c.radius * math.Pi
}
