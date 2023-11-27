package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在原来的内容上追加"ABC/English"
func main() {
	//1.打开一个已经存在的文件
	filePath := "d:/Try.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 666)
	if err != nil {
		fmt.Printf("open file error = %v\n", err)
		return
	}
	//关闭file指针
	defer file.Close()
	//准备写入
	str := "ABC/English\n"
	//写入时使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//因为writer带缓存，所以写入的文本是写在缓存里了，需要进一步提交
	writer.Flush() //将缓存的数据提交到文件中
}
