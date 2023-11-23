package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myInt int

	// 使用 unsafe 包中的 Sizeof 方法获取整型变量占用的字节数
	size := unsafe.Sizeof(myInt)

	// 打印结果
	fmt.Printf("整型变量占用字节数: %d\n", size)
}
