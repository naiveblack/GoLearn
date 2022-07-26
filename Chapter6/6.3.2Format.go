package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// func Itoa(i int) string
	// Itoa 是FormatInt(int64(i), 10)的缩写
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

	// func FormatInt(i int64, base int) string
	// 其中base表示输出的字符串是由十进制转到多少进制
	// >=10的字符输出都是小写字母
	v := int64(1314)
	s10 := strconv.FormatInt(v, 10)
	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s10, s10)
	fmt.Printf("%T, %v\n", s16, s16) //输出的是16进制

	// func FormatUint(i uint64, base int) string 同理FormatInt 无符号

	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	/*
		FormatFloat 根据格式 fmt 和精度 prec 将浮点数 f 转换为字符串。
		假设原始值是从 bitSize 位的浮点值(32 表示 float32，64 表示 float64)获得的，它会对结果进行四舍五入。

		格式fmt是'b'(-ddddp±ddd，二进制 index )、'e'(-d.dddde±dd，十进制 index )、
				'E'(-d.ddddE±dd，十进制 index )之一), 'f' (-ddd.dddd, 无 index ),
				'g' ('e' 用于大 index , 'f' 否则), 'G' ('E' 用于大 index , 'f' 否则),
				'x'(-0xd.ddddp±ddd，十六进制分数和二进制 index )，或'X'(-0Xd.ddddP±ddd，十六进制分数和二进制 index )。

		精度 prec 控制由 'e'、'E'、'f', 'g'、'G'、'x' 和 'X' 格式打印的位数(不包括 index )。
		对于'e'、'E'、'f', 'x'和'X'，它是小数点后的位数。对于'g' 和'G'，它是有效数字的最大数量(删除尾随零)。
		特殊精度 -1 使用所需的最少位数，以便 ParseFloat 将准确返回 f。
	*/
	vf := 3.1415926535
	sf32 := strconv.FormatFloat(vf, 'E', -1, 32)
	fmt.Printf("%T, %v\n", sf32, sf32)
	sf64 := strconv.FormatFloat(vf, 'E', -1, 64)
	fmt.Printf("%T, %v\n", sf64, sf64)

	// func FormatBool(b bool) string
	// 根据b的值返回true或者false
	v1 := true
	s1 := strconv.FormatBool(v1)
	fmt.Printf("%T, %v\n", v1, s1)

	fmt.Println(time.Now())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Hour())
}
