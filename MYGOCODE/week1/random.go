package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate and print 10 random integers using rand.Int()
	for i := 0; i < 10; i++ {
		a := rand.Int() // [1](https://www.geeksforgeeks.org/generating-random-numbers-in-golang/)
		fmt.Printf("%d / ", a)
	}
	fmt.Println("\n----------------------")
	// Generate and print 5 random integers in the range [0, 8) using rand.Intn()
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // [2](https://www.tutorialspoint.com/generating-random-numbers-in-golang)
		fmt.Printf("%d / ", r)
	}

	fmt.Println()

	// Seed the random number generator with the current nanosecond time
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	// Generate and print 10 random float numbers between 0 and 100 using rand.Float32()
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32()) // [6](https://copyprogramming.com/howto/go-golang-get-a-random-number-code-example)
	}
}
