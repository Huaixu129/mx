package main

import "fmt"

func main() {

}

type Pupil struct {
	Age   int
	Name  string
	Score int
}

func (p *Pupil) showInfo() {
	fmt.Printf("name: %v age: %v score: %v\n", p.Name, p.Age, p.Score)
}

func (p *Pupil) SetScore(score int) {
	p.Score = score
}
