package main

import "fmt"

type Phone struct {

}

func (p *Phone) Show(i interface{}) interface{} {
	return i
}

func (p *Phone) Start() {
	fmt.Println("phone is start")
}

func (p *Phone) Stop() {
	fmt.Println("phone is stop")
}

