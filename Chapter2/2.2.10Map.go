package main

import "fmt"

func main() {
	//创建一个map, key，value为 string 类型
	m := make(map[string]string)
	m["username"] = "admin"
	m["sex"] = "man"
	m["age"] = "20"
	fmt.Println(m)
	//删除键值
	delete(m, "age")
	fmt.Println(m)
	//查询键值是否存在
	value, ok := m["username"] // ok判断username是否在map中 value获取key对应的值
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("nil") //输出：如果 ok 返回 false，则输出 nil
	}
}
