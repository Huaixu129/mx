package main

import "fmt"

func main() {
	f := counter()
	fmt.Println(f())
	fmt.Println(f())
}
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
