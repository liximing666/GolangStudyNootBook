package main

import "fmt"

type Phone struct {
	name string
	price int
}

type PhoneMethod interface {
	call()
	photo()
}


type Iphone struct {
	Phone
}
func (i *Iphone) call() {
	fmt.Println("iphone call")
}

func (i *Iphone) photo() {
	fmt.Println("iphone photo")
}


type Sony struct {
	Phone
}
func (s *Sony) call() {
	fmt.Println("sony call")
}
func (s *Sony) photo() {
	fmt.Println("sony photo")
}

func show(inf PhoneMethod)  {
	inf.call()
	inf.photo()
}

func main() {

	var iphone Iphone
	iphone.name = "iphone"
	iphone.price = 6999

	var sony Sony
	sony.name = "sony"
	sony.price = 5999
	//多态性的体现
	show(&iphone)
	fmt.Println("***********************")
	show(&sony)



}