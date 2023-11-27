package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "D:/markdown/线性代数.md"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v", err)
	}
	//显示读取内容
	fmt.Printf("%v", string(content)) //[]byte
	//没有显式的open也不需要显式的close,因为这两个方法都被封装到了readfile函数内部

}
