package main

import "fmt"

func creat_counter() func() int {
	counut := 0
	return func() int {
		counut++
		return counut
	}
}

func X(n int, counter func() int) int {
	count := counter()
	fmt.Println(count)
	if n <= 3 {
		return 1
	}
	return X(n-2, counter) + X(n-4, counter) + 1
}

func main() {
	counter := creat_counter()
	X(8, counter)
}
