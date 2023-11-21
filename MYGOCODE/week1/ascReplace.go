// package main
//
// import (
//
//	"fmt"
//	"strings"
//	"unicode"
//
// )
//
//	func replaceNonASCII(text string) string {
//		// 定义映射函数，将非ASCII字符替换为 ?
//		mapFunc := func(r rune) rune {
//			if r > unicode.MaxASCII {
//				return '?'
//			}
//			return r
//		}
//
//		// 使用 strings.Map 进行替换
//		result := strings.Map(mapFunc, text)
//		return result
//	}
//
//	func main() {
//		// 示例文本
//		inputText := "Hello, 你好，こんにちは"
//
//		// 替换非ASCII字符
//		result := replaceNonASCII(inputText)
//
//		// 打印结果
//		fmt.Println(result)
//	}
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func replaceNoAsc(text string) string {
	mapfunc := func(r rune) rune {
		if r > unicode.MaxASCII {
			return '?'
		}
		return r
	}
	result := strings.Map(mapfunc, text)
	return result
}
func main() {
	inputText := "hello,你好,槐xu"
	result := replaceNoAsc(inputText)
	fmt.Println(result)
}
