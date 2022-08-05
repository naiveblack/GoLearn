package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}
type Student struct {
	Person     //采用匿名结构体字段模拟继承关系
	SchoolName string
}

func main() {
	//初始化 Person
	p1 := Person{"Steven", 35, "男"}
	fmt.Println(p1)
	fmt.Printf("p1: %T , %+v \n", p1, p1)
	fmt.Println("----------------------")

	//初始化 Student

	//写法 1：
	s1 := Student{p1, "北航软件学院"}
	fmt.Println(s1)
	fmt.Printf("s1: %T , %+v \n", s1, s1)
	fmt.Println("----------------------")

	//写法 2：
	s2 := Student{Person{"Josh", 30, "男"}, "北外高翻学院"}
	fmt.Println(s2)
	fmt.Printf("s2: %T , %+v \n", s2, s2)
	fmt.Println("----------------------")

	//写法 3：
	s3 := Student{
		Person: Person{
			Name: "Penn",
			Age:  19,
			Sex:  "男",
		},
		SchoolName: "北大元培学院",
	}
	fmt.Println(s3)
	fmt.Printf("s3: %T , %+v \n", s3, s3)
	fmt.Println("----------------------")

	//写法 4：
	s4 := Student{}
	s4.Name = "Daniel"
	s4.Sex = "男"
	s4.Age = 12
	s4.SchoolName = "十一龙樾"
	fmt.Println(s4)
	fmt.Printf("s4: %T , %+v \n", s4, s4)
	fmt.Println("----------------------")
}
