package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		x, y := 1, 1
		for i := 0; i < 10; i++ {
			tmp := x + y
			ch <- x
			x = y
			y = tmp
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()

	for {
		select {
		case num := <- ch:
			fmt.Println(num)
		case <- quit:
			return
		}
	}
}
