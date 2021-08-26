package main

import (
	"fmt"
	"runtime"
	"time"
)

//select每一个case 只能判断io select中的 break 退出的是当前的 select

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()

	for {
		select {
			case n := <- ch:
				fmt.Println(n)
			case <- quit:
				return // 用return退出main
		}
		fmt.Println("----------------------------------")

	}

}
