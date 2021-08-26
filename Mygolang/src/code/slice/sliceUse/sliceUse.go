package main

import "fmt"



func main() {
	fmt.Println("slice 的声明方式")

	slice1 := []int {1,2,3}
	println(slice1)

	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Printf("len = %d , slice2 = %v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d , slice3 = %v\n", len(slice3), slice3)

	slice4 := make([]int, 3)
	fmt.Printf("len = %d , slice4 = %v\n", len(slice4), slice4)

	var slice5 []int
	println(slice5 == nil)


}
