package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("the %dth parameter's type is bool,value is %v\n", index, x)
		case float64, float32:
			fmt.Printf("the %dth parameter's type is float,value is %v\n", index, x)
		case int, int32, int64:
			fmt.Printf("the %dth parameter's type is int,value is %v\n", index, x)
		case string:
			fmt.Printf("the %dth parameter's type is string,value is %v\n", index, x)
		default:
			fmt.Println("the parameter's type isn't definate")
		}
	}
	prt(&student{})
}
func main() {
	var n1 float32 = 32.0
	var n2 float64 = 64.0
	var n3 int = 32
	var n4 string = "hello"
	var n5 bool = false
	n6 := 300
	n7 := "world"
	TypeJudge(n1, n2, n3, n4, n5, n6, n7)
}

type student struct {
}

func prt(a *student) {
	fmt.Printf("%T", a)
}
