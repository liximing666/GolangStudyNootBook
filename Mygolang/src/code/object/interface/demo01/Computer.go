package main

import "fmt"

type Computer struct {

}

func (c *Computer) Show(i interface{}) interface{} {
	return i
}

func (c *Computer) Start() {
	fmt.Println("computer is start")
}

func (c *Computer) Stop() {
	fmt.Println("computer is stop")
}

