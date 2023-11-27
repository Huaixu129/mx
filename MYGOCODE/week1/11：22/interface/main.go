package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Cmera struct {
}

func (p Phone) Start() {
	fmt.Println("phone start working...")
}

func (p Phone) Stop() {
	fmt.Println("phone stop working...")
}

func (p Phone) Call() {
	fmt.Println("phone start calling...")
}
func (c Cmera) Start() {
	fmt.Println("cmera strat working...")
}

func (c Cmera) Stop() {
	fmt.Println("cmera stop working...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	//fmt.Println("is working")
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	cp := Computer{}
	ph := Phone{}
	cm := Cmera{}
	cp.Working(ph)
	cp.Working(cm)
}
