package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n==== ==== ==== ====")

	// 初试语句init
	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n==== ==== ==== ====")

	//条件表达式condition
	i = 0
	for ; ; i++ {
		if i > 20 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n==== ==== ==== ====")

	//只有一个条件表达式 效果类似while循环
	i = 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("\n==== ==== ==== ====")

	// 无表达式 与 for ( ; ; )一致
	i = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("\n==== ==== ==== ====")

	//for range 迭代循环（字符串、slice、数组、map
	str := "123ABCabc一丁丂"
	for j, value := range str {
		fmt.Printf("第%d位的ASCII值=%d, 字符是%c\n", j, value, value)
	}
	fmt.Println("\n==== ==== ==== ====")
}
