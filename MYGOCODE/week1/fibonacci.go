package main

import "fmt"

func main() {
	num1 := getFbonacci(6)
	fmt.Println(num1)
	zton(10)
	fmt.Println("--------------------------------------")
	fmt.Println(factorial(3))
}

func getFbonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return getFbonacci(n-1) + getFbonacci(n-2)
}
func zton(n int) {
	fmt.Println(n)
	if n == 1 {
		return
	}
	zton(n - 1)
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
