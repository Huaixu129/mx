package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("所有命令行参数值有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%v\n\r", i, v)
	}
}
