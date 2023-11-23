package main

import (
	"MYGOCODE/week1/11：22/encapsulation/demo"
	"fmt"
)

func main() {
	p1 := demo.NewPerson("huaixu")
	fmt.Println(p1)
	p1.SetAge(21)
	p1.SetSallary(30000)
	fmt.Println(p1)
	fmt.Println(p1.Name, "age = ", p1.GetAge(), "salary = ", p1.Getsalay())
}
