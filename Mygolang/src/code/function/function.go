package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

func main() {
	res := add(3, 4)
	fmt.Println(res)

}
