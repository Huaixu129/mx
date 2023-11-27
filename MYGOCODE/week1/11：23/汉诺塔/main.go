package main

import "fmt"

func Move(disk int, source, target string) {
	fmt.Printf("Move disk %d from %s to %s\n", disk, source, target)
}

func Hanoi(n int, source, auxiliary, target string) {
	if n == 1 {
		Move(1, source, target)
	} else {
		Hanoi(n-1, source, target, auxiliary)
		Move(n, source, target)
		Hanoi(n-1, auxiliary, source, target)
	}
}

func main() {
	n := 2
	Hanoi(n, "A", "B", "C")
}
