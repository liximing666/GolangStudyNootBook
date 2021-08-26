package main

import "fmt"

type Send struct {
	values []int
	re *Recive
}

func (send *Send) SendEle() {
	for i := 0; i < 3; i++ {
		send.re.recives <- i
	}
}

type Recive struct {
	recives chan int
}
func (recive *Recive) Listen() {
	for  {
		x := <- recive.recives
		fmt.Println(x)
	}
}

func NewSend() *Send {
	return &Send{values: []int{1,2,3}, re: &Recive{}}
}

func NewRecive() *Recive {
	res := &Recive{recives: make(chan int)}
	res.Listen()
	return res
}

func main() {

	r := NewRecive()
	s := NewSend()
	s.re = r

	go s.SendEle()

	for  {
		;
	}


}
