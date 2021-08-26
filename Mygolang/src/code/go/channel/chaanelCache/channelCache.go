package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("write", i)
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 6; i++ {
		num := <- c
		fmt.Println("read", num)
	}

	close(c)

}
