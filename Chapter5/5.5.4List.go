package main

import (
	"container/list"
	"fmt"
)

func main() {
	//定义list
	// l := list.New()
	var l list.List

	// 添加数据 类似于链表 但是go语言中list可以添加不同类型的元素
	l.PushBack(1)
	l.PushBack("aaa")
	l.PushBack(5.67)

	//正序输出
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//倒序数据
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	fmt.Println(l)

	copyList()
}

// list是值类型 通过var命名则不会相互影响
// 但通过New()方法声明则是一个指针 在拷贝和参数传递时具有引用类型的特征
func copyList() {
	//声明 list1
	//var list1 list.List
	list1 := list.New()
	printListInfo2("刚声明的 list1：", list1)
	//给 list1 赋值
	list1.PushBack("one")
	list1.PushBack(2)
	list1.PushBack("three")
	list1.PushFront("first")
	printListInfo2("赋值后的 list1", list1)
	iterateList2(list1)
	//将 list1 拷贝给 list2。其实拷贝的是地址
	list2 := list1
	printListInfo2("刚拷贝的 list2", list2)
	iterateList2(list2)
	//list2 修改后
	list2.PushBack(250)
	list2.PushBack(350)
	list2.PushBack(450)
	printListInfo2("修改后的 list2", list2)
	iterateList2(list2)
	//list2 的修改是否影响到 list1？
	printListInfo2("修改 list2 的 list1", list1)
	iterateList2(list1)
}

func printListInfo2(info string, l *list.List) {
	fmt.Println(info + "----------")
	fmt.Printf("%T:%v \t ， 长度为：%d \n", l, l, l.Len())
	fmt.Println("----------")
}

func iterateList2(l *list.List) {
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		fmt.Printf("%d:%v \t", i, e.Value)
	}
	fmt.Println("\n----------")
}
