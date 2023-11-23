package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Alice", Age: 30}

	// 使用 reflect.DeepEqual 比较结构体是否相等
	if reflect.DeepEqual(person1, person2) {
		fmt.Println("两个结构体相等")
	} else {
		fmt.Println("两个结构体不相等")
	}
}
