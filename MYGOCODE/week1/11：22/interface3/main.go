package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Age  int
	Name string
}

type herolist []Hero

func (hl herolist) Len() int {
	return len(hl)
}

func (hl herolist) Less(i, j int) bool {
	return hl[i].Age < hl[j].Age
}
func (hl herolist) Swap(i, j int) {
	temp := hl[i]
	hl[i] = hl[j]
	hl[j] = temp
}
func main() {
	var heros herolist
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(50)),
			Age:  rand.Intn(30),
		}
		heros = append(heros, hero)
	}
	for _, v := range heros {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%v", v)
	}
}
