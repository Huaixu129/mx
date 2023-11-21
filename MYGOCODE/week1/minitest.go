package main

import "fmt"

//	func f() (ret int) {
//		defer func() {
//			ret++
//		}()
//		return 1
//	}
func main() {
	//fmt.Println(f())
	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s1, s2)
}
