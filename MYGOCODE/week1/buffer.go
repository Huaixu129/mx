package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer // 1. 初始化一个 bytes.Buffer

	// 2. 循环获取字符串，将其写入 buffer
	for i := 1; i <= 3; i++ {
		s := fmt.Sprintf("String %d", i)
		buffer.WriteString(s)
	}

	// 3. 打印串联后的字符串
	fmt.Println(buffer.String()) // 输出：String 1String 2String 3
}
