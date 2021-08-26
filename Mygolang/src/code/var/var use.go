package main

import "fmt"

var a = 100

func main() {
	age := 10
	fmt.Println(age)

	const (
		a = iota
		b
		c
		d
		e
	)


	fmt.Println(e)
}
