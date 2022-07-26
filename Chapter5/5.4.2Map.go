package main

import "fmt"

func main() {
	// 定义一个map
	var country = map[string]string{
		"China":  "Beijing",
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}
	fmt.Println(country)

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Println(rating)

	countryCapitalMap := make(map[string]string)
	fmt.Println(countryCapitalMap)
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	fmt.Println(countryCapitalMap)

	// 遍历map
	for k, v := range countryCapitalMap {
		fmt.Println("国家: ", k, " 首都: ", v)
	}

	// 利用map查看元素是否存在
	// ok返回的是对应key的元素 如果key不存在 则返回默认值0/false 不会报错
	captial, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("首都为: ", captial)
	} else {
		fmt.Println("该国家不在map中")
	}
}
