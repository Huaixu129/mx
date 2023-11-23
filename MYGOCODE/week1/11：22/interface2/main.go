package main

import (
	"fmt"
)

type Bitfc interface {
	test01()
}
type Citfc interface {
	test02()
}
type Aitfc interface {
	Bitfc
	Citfc
	test03()
}
type Stu struct {
}

func (s Stu) test01() {
	fmt.Println("hello")
}
func (s Stu) test02() {
	fmt.Println("hi")
}
func (s Stu) test03() {
	fmt.Println("how are you")
}
func main() {
	var student Stu
	var A Aitfc = student
	//var B Bitfc = student
	//var C Citfc = student
	A.test01()
	A.test02()
	A.test03()
}
