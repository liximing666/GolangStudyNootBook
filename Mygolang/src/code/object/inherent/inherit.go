package main

import "fmt"

type Animal struct {
	name string
}
func (this *Animal) Eat() {
	fmt.Println("annimal eat")
}

type Cat struct {
	Animal
	hobby string
}

func (this *Cat) Eat() {
	fmt.Println("cat eat")
}



func main() {

	a := Animal{name: "tom"}
	a.Eat()

	var cat Cat
	cat.name = "miao"
	cat.hobby = "run"
	cat.Eat()

}
