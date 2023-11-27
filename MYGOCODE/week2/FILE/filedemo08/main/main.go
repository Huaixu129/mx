package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	chrCount int
	numCount int
	spcCount int
	otrCount int
}

func main() {
	var count charCount
	//创建一个reader
	file, err := os.Open("d:/Try.txt")
	if err != nil {
		fmt.Printf("file open error = %v", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.chrCount++
			case v >= '0' && v <= '9':
				count.numCount++
			case v == ' ' || v == '\t':
				count.spcCount++
			default:
				count.otrCount++
			}
		}
	}
	//输出结果
	fmt.Println(count)
}
