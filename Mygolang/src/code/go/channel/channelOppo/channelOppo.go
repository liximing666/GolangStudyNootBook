package main

import (
	"fmt"
	"time"
)

//利用管道实现互斥

func Printer(s string)  {
	for _, chat := range s{
		fmt.Printf("%c", chat)
		time.Sleep(300 * time.Microsecond)
	}
}

func main() {

	c := make(chan int, 1)
	c <- 6

	go func() {
		<- c
		Printer("hello")
		c <- 6
	}()

	go func() {
		<- c
		Printer("world")
		c <- 6
	}()

	for  {
		;
	}
}
