package main

import "fmt"

func reverseStringWithByteSlice(s string) string {
	// 将字符串转换为 []byte 类型的切片
	bytes := []byte(s)
	length := len(bytes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
func reverseStringWithSingleSlice(s string) string {

	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
func main() {
	input := "Google"

	resultByteSlice := reverseStringWithByteSlice(input)
	fmt.Println("反转字符串（[]byte 类型切片）:", resultByteSlice)

	resultSingleSlice := reverseStringWithSingleSlice(input)
	fmt.Println("反转字符串（单个切片交换法）:", resultSingleSlice)
}
