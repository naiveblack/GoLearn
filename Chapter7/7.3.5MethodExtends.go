package main

import "fmt"

func main() {
	s1 := Student1{Human{"David", "111222333", 13}, "**学校"}
	e1 := Employee1{Human{"Caster", "8888", 28}, "**公司"}
	fmt.Println(s1)
	fmt.Println(e1)
}

type Human struct {
	name, phone string
	age         int
}

type Student1 struct {
	Human
	school string
}

type Employee1 struct {
	Human
	company string
}
