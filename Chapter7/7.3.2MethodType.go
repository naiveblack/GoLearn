package main

import "fmt"

type Employee struct {
	name, currency string
	salary         int
}

func main() {
	emp1 := Employee{
		name:     "David",
		salary:   2000,
		currency: "$",
	}

	//调用方法
	emp1.displaySalary()

	//调用函数
	displaySalary(emp1)
}

func (e Employee) displaySalary() {
	fmt.Printf("员工姓名： %s , 薪资： %s%d\n", e.name, e.currency, e.salary)
}

func displaySalary(e Employee) {
	fmt.Printf("员工姓名： %s , 薪资： %s%d\n", e.name, e.currency, e.salary)
}
