package main

import "fmt"

func main() {
	var str string = "hello,世界"
	fmt.Println(len(str))
	for i, j := range str {
		fmt.Println(i, string(j))
	}
}
