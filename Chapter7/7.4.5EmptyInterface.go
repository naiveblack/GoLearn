package main

import "fmt"

//定义空接口
type A interface {
}
type Cat struct {
	name string
	age  int
}
type Personp struct {
	name string
	sex  string
}

func main() {
	// 用空接口表示任意数据类型。
	var a1 A = Cat{name: "Mimi", age: 1}
	var a2 A = Personp{"Steven", "man"}
	var a3 A = "Learn golang with me!"
	var a4 A = 100
	var a5 A = 3.14
	fmt.Printf("%T，%v \n", a1, a1)
	fmt.Printf("%T，%v \n", a2, a2)
	fmt.Printf("%T，%v \n", a3, a3)
	fmt.Printf("%T，%v \n", a4, a4)
	fmt.Printf("%T，%v \n", a5, a5)
	fmt.Println("------------------------")
	//println 的参数就是空接口
	fmt.Println("println 的参数可以是任何数据类型，用空接口表示\n", 100, 3.14, Cat{"小天", 2})
	//定义一个 map：key 是 string，value 是任意数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	fmt.Println(map1)
	fmt.Println("------------------------")
	//定义一个切片，其中存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
	testInterface(slice1)
}
func testInterface(s []interface{}) {
	for i := range s {
		fmt.Println("第", i+1, "个数据：")
		switch ins := s[i].(type) {
		case Cat:
			fmt.Println("\tcat 对象：", ins.name, ins.age)
		case Personp:
			fmt.Println("\tperson 对象：", ins.name, ins.sex)
		case int:
			fmt.Println("\tint 类型：", ins)
		case string:
			fmt.Println("\tstring 类型：", ins)
		case float64:
			fmt.Println("\tfloat64 类型：", ins)
		}
	}
}
