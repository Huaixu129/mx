package main

import "fmt"

type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("hello")
}

type integer int

func (i integer) Say() {
	fmt.Println("sayi")
}
func main() {
	var student Stu
	var a Ainterface = student
	a.Say()
	var i integer
	var b Ainterface = i
	b.Say()
}
