package main

import "fmt"

type Call interface {
	Call()
	Show()
}

type Phone struct {
	badge string
}

func (p *Phone) Call() {
	fmt.Println("im call")
}

func (p *Phone) Show() {
	fmt.Println("im show")
}

func main() {
	p := Phone{badge: "apple"}
	p.Call()
	p.Show()

//接口是个指针 本质是在标准化
	var callInf Call
	callInf = &Phone{badge: "a"}
	callInf.Call()
	callInf.Show()


//多态性的具体例子
}
