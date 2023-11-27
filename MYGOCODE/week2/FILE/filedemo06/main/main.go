package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//将d:/try.txt中的内容导入到e:/trytry.text
	file1Path := "d:/Try.txt"
	file2Path := "e:/TryTry.txt"
	date, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file nil = %v", err)
		return
	}
	err = ioutil.WriteFile(file2Path, date, 666)
	if err != nil {
		fmt.Printf("write file nil = %v", err)
		return
	}
	fmt.Println(PathExists(file1Path))
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
