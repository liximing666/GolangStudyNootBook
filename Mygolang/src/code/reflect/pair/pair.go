package main

import "fmt"

type readbook interface {
	Readbook()
}
type writebook interface {
	Writebook()
}

type Book struct {

}

func (b *Book) Readbook() {
	fmt.Println("read book")
}

func (b *Book) Writebook() {
	fmt.Println("write book")
}

func main() {
	b := &Book{}
	//pair : value:b type:Book
	var r readbook = b
	var w writebook = b

	if w == nil {}

	_, ok := r.(writebook)
	fmt.Println(ok)
	//这里readbook 和 writebook判定为一种类型是因为都是从b pair里的type传递过来的
 	


}
