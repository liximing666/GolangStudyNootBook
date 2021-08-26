package main

import "fmt"

//接口之间是可以继承的
type I3 interface {
	Show3()
}

type I2 interface {
	Show2()
}

type I1 interface {
	Show1()
	I3
	I2
}

type Show struct {

}

func (s Show) Show1() {
	fmt.Println("1")
}

func (s Show) Show3() {
	fmt.Println("3")}

func (s Show) Show2() {
	fmt.Println("2")
}

