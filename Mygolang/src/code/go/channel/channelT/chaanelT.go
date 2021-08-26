package main

import "fmt"

func main() {

	v := make(chan int, 8)
	mutex := make(chan int)

	go func() {

		for i := 0; i < 8; i++ {
			v <- i
		}
		close(v)
		mutex <- 6
	}()

	<- mutex

	for value := range v {

		fmt.Println(value)
	}

}
