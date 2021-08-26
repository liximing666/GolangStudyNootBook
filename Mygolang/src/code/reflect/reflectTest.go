package main

import "fmt"

//反射依托于断言机制，通过断言拿到 <type value> 这个pair ，在赋值的时候这个pair会继续传递下去

func main() {
	var A string
	A = "abcde"

	var B interface{} = A

	_, ok := B.(string)
	fmt.Println(ok)

}
