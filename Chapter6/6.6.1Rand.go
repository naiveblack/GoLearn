package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成固定的随机数 默认情况下 随机数种子都是1 seed是64位整数
	a := rand.Int()
	b := rand.Float64()
	c := rand.Intn(10) //从0-10之间取随机数
	fmt.Println(a, b, c)

	// 动态随机数种子
	// UnixNano 将 t 表示为 Unix 时间，即从时间点 January 1, 1970 UTC 到时间点 t 所经过的时间（单位纳秒）
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randNum := r1.Intn(100)
	fmt.Println(randNum)

	// 简写形式
	// 1 获取随机0-10的整数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))

	// 2 获取0.0 - 1.0之间的随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Float64())

	// 3 获取两数之间的随机数	 10--20之间
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10) + 10)
}
