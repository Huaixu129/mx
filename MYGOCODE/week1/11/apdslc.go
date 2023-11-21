// package main
//
// import "fmt"
//
//	func main() {
//		fmt.Println()
//
// }
//
//	func Append(slice []int, date ...int) []int {
//		m := len(slice)
//		n := m + len(date)
//		if n > cap(slice) {
//			newslice := make([]int, (n+1)*2)
//			copy(newslice, slice)
//			slice = newslice
//		}
//		slice = slice[0:n]
//		copy(slice[m:n], date)
//		return slice
//	}
package main

import "fmt"

func extendSlice(s []int, factor int) []int {
	// 计算新的长度
	newLength := len(s) * factor

	// 使用append扩展切片
	s = append(s, make([]int, newLength-len(s))...)

	return s
}

func main() {
	// 示例用法
	originalSlice := []int{1, 2, 3}
	factor := 3

	extendedSlice := extendSlice(originalSlice, factor)

	fmt.Println(extendedSlice)
}
