package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.func Contains(s, substr string) bool
	//判断字符串 s 是否包含 substr 字符串
	fmt.Println(strings.Contains("seafood", "foo")) //输出：true
	fmt.Println(strings.Contains("seafood", "bar")) //输出：false
	fmt.Println(strings.Contains("seafood", ""))    //输出：true
	fmt.Println(strings.Contains("", ""))           //输出：true

	//2.func ContainsAny(s, chars string) bool
	//判断字符串 s 是否包含 chars 字符串中的任一字符
	fmt.Println(strings.ContainsAny("team", "i"))        //输出：false
	fmt.Println(strings.ContainsAny("failure", "u & i")) //输出：true
	fmt.Println(strings.ContainsAny("foo", ""))          //输出：false
	fmt.Println(strings.ContainsAny("", ""))             //输出：false

	//3.func Count(s, sep string) int
	//返回字符串 s 包含字符串 sep 的个数
	//如果 sep 是一个空切片，则 Count返回 1+s 中 UTF-8 编码的代码点数。
	fmt.Println(strings.Count("cheese", "e")) //输出：3
	fmt.Println(strings.Count("five", ""))    //输出：5

	//4.func Index(s, sep string) int
	//返回字符串 s 中字符串 sep 首次出现的位置
	fmt.Println(strings.Index("chicken", "ken")) //输出：4
	fmt.Println(strings.Index("chicken", "dmr")) //输出：-1

	//5.func IndexAny(s, chars string) int
	//返回字符串 chars 中的任一 unicode 码值 r 在 s 中首次出现的位置
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) //输出：2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   //输出：-1

	//6.func IndexFunc(s string, f func(rune) bool) int
	//返回字符串 s 中满足函数 f(r)==true 字符首次出现的位置
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    //输出：7
	fmt.Println(strings.IndexFunc("Hello, world", f)) //输出：-1

	//7.func IndexRune(s string, r rune) int
	//返回 unicode 码值 r 在字符串中首次出现的位置
	fmt.Println(strings.IndexRune("chicken", 'k')) //输出：4
	fmt.Println(strings.IndexRune("chicken", 'd')) //输出：-1

	//8.func LastIndex(s, sep string) int
	//返回字符串 s 中字符串 sep 最后一次出现的位置
	fmt.Println(strings.Index("go gopher", "go"))         //输出：0
	fmt.Println(strings.LastIndex("go gopher", "go"))     //输出：3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) //输出：-1
}
