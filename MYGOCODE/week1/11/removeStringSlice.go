package main

import "fmt"

// RemoveStringSlice 从切片中移除从 start 到 end 索引的元素
func RemoveStringSlice(slice []string, start, end int) []string {
	// 确保 start 和 end 在合法范围内
	if start < 0 || start >= len(slice) || end < start || end >= len(slice) {
		fmt.Println("Invalid start or end index.")
		return slice
	}

	// 使用切片操作将指定范围的元素移除
	slice = append(slice[:start], slice[end+1:]...)

	return slice
}

func main() {
	// 示例用法
	slice := []string{"apple", "banana", "cherry", "date", "fig"}
	startIndex := 1
	endIndex := 3

	result := RemoveStringSlice(slice, startIndex, endIndex)
	fmt.Println(result)
}
