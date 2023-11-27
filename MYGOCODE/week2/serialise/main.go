package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Salary   float64
	Birthday string
	Id       []int
}

// serialise the struct
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      3,
		Salary:   100000,
		Birthday: "2023/11/27",
		Id:       []int{1, 2, 3, 4},
	}
	date, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("serial number error = %v", err)
	}
	fmt.Printf("serialised monster = %v\n", string(date))
}

// serialise the map
func testMap() {
	a := make(map[string]interface{})
	a["name"] = "雪花钱"
	a["age"] = 18
	a["university"] = "华中师范大学"
	date, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("serial number error = %v", err)
	}
	fmt.Printf("serialised monster = %v\n", string(date))
}
func testSlice() {
	s := new([]map[string]interface{})
	m1 := make(map[string]interface{})
	m1["name"] = "雪花钱"
	m1["age"] = 18
	m1["university"] = "华中师范大学"
	m2 := make(map[string]interface{})
	m2["name"] = "槐序"
	m2["age"] = 18
	m2["university"] = "华中师范大学"
	*s = append(*s, m1)
	*s = append(*s, m2)
	date, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("serial number error = %v", err)
	}
	fmt.Printf("serialised monster = %v\n", string(date))
}

// 基本数据类型序列化
func float64Test() {
	var f float64 = 123.456
	date, err := json.Marshal(f)
	if err != nil {
		fmt.Printf("serial number error = %v", err)
	}
	fmt.Printf("serialised float64 = %v\n", string(date))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	float64Test()
}
