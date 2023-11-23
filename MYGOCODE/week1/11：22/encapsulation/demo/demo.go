package demo

import "fmt"

// 封装，公开->大写，隐藏->小写(其他包不能直接访问)
type person struct {
	Name   string
	age    int
	salary float64
}

// 写一个工厂函数来访问公开数据
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
		return
	}
	fmt.Println("年龄范围不合理")
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSallary(sal float64) {
	if sal > 3000 && sal <= 30000 {
		p.salary = sal
		return
	}
	fmt.Println("薪水范围不合理")
}

func (p *person) Getsalay() float64 {
	return p.salary
}
