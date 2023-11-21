package main

import "fmt"

// InsertStringSlice 将切片插入到另一个切片的指定位置
func InsertStringSlice(dest, src []string, index int) []string {
	// 创建一个临时切片用于存储插入位置之后的元素
	temp := append([]string{}, dest[index:]...)

	// 将源切片插入到目标切片的指定位置
	dest = append(dest[:index], src...)
	dest = append(dest, temp...)

	return dest
}

func main() {
	// 示例用法
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"orange", "grape"}
	index := 1

	result := InsertStringSlice(slice1, slice2, index)
	fmt.Println(result)
}
