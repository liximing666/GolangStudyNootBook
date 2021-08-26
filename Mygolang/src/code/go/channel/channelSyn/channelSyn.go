package main

//用管道来实现两个进程的同步关系 （pv操作信号量为 1）

import (
	"fmt"
	"runtime"
	"time"
)

func Printer(s string)  {
	for _, chat := range s{
		fmt.Printf("%c", chat)
		time.Sleep(300 * time.Microsecond)
	}
}


func main() {

	c := make(chan int)

	runtime.GOMAXPROCS(1)

	go func() {
		Printer("hello")
		c <- 6

	}()

	go func() {
		<- c
		Printer("world")
	}()

	for  {
		;
	}

}
