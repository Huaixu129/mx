package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 拷贝文件
// 编写一个函数接收文件路径
func copyFile(dst string, src string) (written int64, err error) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file error = %v", err)
	}
	//通过srcfile获取到file指针
	reader := bufio.NewReader(srcfile)
	defer srcfile.Close()
	//打开dst
	dstfil, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 666)
	if err != nil {
		fmt.Printf("write file error = %v", err)
	}
	writer := bufio.NewWriter(dstfil)
	defer dstfil.Close()
	return io.Copy(writer, reader)
}
func main() {
	src := "E:/《仙剑问情女神》画师：炼丹 高清无水印原图4K壁纸_彼岸图网(1).jpg"
	dst := "F:/abc.jpg"
	_, err := copyFile(dst, src)
	if err == nil {
		fmt.Println("copy complete")
	} else {
		fmt.Printf("copyflie error = %v", err)
	}
}
