package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	//类型断言
	var b Point
	b = a.(Point)
	fmt.Println(b)

	var t float64 = 1.1
	var d interface{}
	d = t
	var c float64
	c = d.(float64)
	fmt.Println(c)
}
