package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个新的双向链表
	myList := list.New()

	// 将 101、102 和 103 放入链表中
	myList.PushBack(101)
	myList.PushBack(102)
	myList.PushBack(103)

	// 打印链表中的元素
	fmt.Print("Elements in the list: ")
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%v ", element.Value)
	}
	fmt.Println()
}
