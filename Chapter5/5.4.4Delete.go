package main

import "fmt"

func main() {
	map1 := map[string]string{
		"element":    "div",
		"width":      "100px",
		"height":     "200px",
		"border":     "solid",
		"color":      "red",
		"background": "none",
	}
	fmt.Println("删除前：", map1)
	if _, ok := map1["background"]; ok {
		delete(map1, "background")
	}
	fmt.Println("删除后：", map1)
	// 重新make一个新的map清空所有数据
	map1 = make(map[string]string)
	fmt.Println("清空后：", map1, len(map1))

	// map与slice相同 都是引用类型
	// 因此一个map给一个新的变量赋值时 其指针将会指向相同的地址
	personSalary := map[string]int{
		"Steven": 18000,
		"Daniel": 5000,
		"Josh":   20000,
	}
	fmt.Println("原始薪资：", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["Daniel"] = 8000
	fmt.Println("修改后 newPersonSalary：", newPersonSalary)
	fmt.Println("personSalary 受影响情况：", personSalary)
}
