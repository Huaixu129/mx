package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world, and hello huaixu"
	p := strings.NewReader(str)
	fmt.Println(*p)
}
