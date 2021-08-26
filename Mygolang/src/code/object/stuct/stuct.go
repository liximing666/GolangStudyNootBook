package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	gender bool
}

func Change(p *Person)  {
	p.age = 16
	p.gender = true
	p.name = "lucy"
}


func NewPerson(name string, age int, gender bool) *Person {
	return &Person{
		name: name,
		age: age,
		gender: gender,
	}
}

func main() {

	person1 := Person{name: "a", age: 3, gender: true}
	fmt.Println(person1.name)

	person2 := Person{}
	person2.name = "b"
	person2.age = 16
	person2.gender = true
	fmt.Println(person2.age)

	fmt.Println(person1)
	Change(&person1)
	fmt.Println(person1)

	pa := NewPerson("aaa", 18, true)
	fmt.Println(pa.name)







}
