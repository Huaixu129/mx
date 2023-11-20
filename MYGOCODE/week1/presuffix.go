package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//判断包含关系
	fmt.Printf("%t\n", strings.Contains(str, "example"))

}
