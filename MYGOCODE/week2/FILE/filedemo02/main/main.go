package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//file 被称为文件对象(or指针or句柄)
	file, err := os.Open("D:/markdown/线性代数.md")
	if err != nil {
		fmt.Println("open file error = ", err)
	}
	defer file.Close() //及时关闭file指针，否则会有句柄泄露
	//创建一个 * Reader,带缓冲,一次一次读，适用于文件较大的情况
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') //读到换行结束
		if err == io.EOF {                  //io.EOF表示读到文件末尾，若没读到则反复读取
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
