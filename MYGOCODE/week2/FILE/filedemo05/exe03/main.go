package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读与写
func main() {
	//1.打开一个已经存在的文件
	filePath := "d:/Try.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 666)
	if err != nil {
		fmt.Printf("open file error = %v\n", err)
		return
	}
	//关闭file指针
	defer file.Close()
	////先读取原来的内容
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	//显示到终端
	//	fmt.Print(str)
	//}
	//准备写入
	str := "hello,beijing\n"
	//写入时使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer带缓存，所以写入的文本是写在缓存里了，需要进一步提交
	writer.Flush() //将缓存的数据提交到文件中
	// 将文件指针重新定位到文件开头
	file.Seek(0, 0)

	// 先读取原来的内容
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 显示到终端
		fmt.Print(str)
	}

}
