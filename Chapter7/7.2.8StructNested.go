package main

import (
	"fmt"
)

type Address struct {
	province, city string
}
type Person1 struct {
	name    string
	age     int
	address Address
}

func main() {
	//模拟对象之间的聚合关系
	p := Person1{}
	p.name = "Steven"
	p.age = 35
	//赋值方式 1
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = addr
	fmt.Println(p)
	fmt.Println("姓名:", p.name)
	fmt.Println("年龄:", p.age)
	fmt.Println("省:", p.address.province)
	fmt.Println("市:", p.address.city)
	fmt.Println("---------------------")
	//修改 Person 对象的数据，是否影响 Address 对象？
	p.address.city = "昌平区"
	fmt.Println("姓名:", p.name)
	fmt.Println("年龄:", p.age)
	fmt.Println("省:", p.address.province)
	fmt.Println("市:", p.address.city)
	fmt.Println("---------------------")
	fmt.Println("市:", addr.city) //没有影响
	fmt.Println("---------------------")

	//修改 Address 对象的数据，是否影响 Person 对象？
	addr.city = "大兴区"
	fmt.Println("姓名:", p.name)
	fmt.Println("年龄:", p.age)
	fmt.Println("省:", p.address.province)
	fmt.Println("市:", p.address.city) //没有影响
	fmt.Println("---------------------")
	//赋值方式 2：
	p.address = Address{
		province: "陕西省",
		city:     "西安市",
	}
	fmt.Println(p)
	fmt.Println("姓名:", p.name)
	fmt.Println("年龄:", p.age)
	fmt.Println("省:", p.address.province)
	fmt.Println("市:", p.address.city)
	fmt.Println("---------------------")
}
