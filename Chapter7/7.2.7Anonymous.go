package main

import (
	"fmt"
	"math"
)

func main() {
	//匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)
	fmt.Println(res)

	// 匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}
	fmt.Println(addr)

	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}
	fmt.Println(cat)

	// 结构体字段匿名 没有字段名的类型 默认以类型为字段名 同种类型只能写一个
	user := User{"Steven", 'm', 25, 176.3}
	fmt.Println(user, user.string)
}

type User struct {
	//name string
	//sex byte
	//age int8
	//height float64
	string
	byte
	int8
	float64
}
