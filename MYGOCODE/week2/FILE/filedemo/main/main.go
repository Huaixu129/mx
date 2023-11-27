package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//file 被称为文件对象(or指针or句柄)
	file, err := os.Open("D:/markdown/线性代数.md")
	if err != nil {
		fmt.Println("open file error = ", err)
	}
	//输出以下文件
	//在这可以看出file就是一个指针
	fmt.Printf("file = %v\n", file)
	//关闭文件
	file.Close()
}
