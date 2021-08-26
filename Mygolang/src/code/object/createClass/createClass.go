package main

import "fmt"

type Person struct {
	name string
	age int
	gender bool
}

//给一个结构体绑定方法，这样属性方法都有了
//绑定结构体传指针才是一个对象

func (this *Person) showInf() {
	fmt.Printf("name = %s, age = %d, gender = %v\n", this.name, this.age, this.gender)
}

func (this *Person) GetName() string {
	return this.name
}

func (this *Person) SetName(name string) {
	this.name = name
}




func main() {
	p := Person{name: "a", age: 16, gender: true}
	p.showInf()
	n := p.GetName()
	fmt.Println(n)
	p.SetName("aaa")
	fmt.Println(p.GetName())


}
