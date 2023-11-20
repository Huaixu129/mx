package main

import (
	"fmt"
)

func main() {
	str1, str2 := "hello", "world"
	F1(str1, str2)
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	fmt.Println(s)
}
func F3(s []string) {
	fmt.Println(s)
}
