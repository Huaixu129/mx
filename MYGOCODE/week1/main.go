package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4}
	for i := range s {
		fmt.Println(i)
	}
}
