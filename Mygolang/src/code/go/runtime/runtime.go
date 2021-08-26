package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("im go rutin")
			//当前让出cpu的时间片
			runtime.Gosched()
		}
		runtime.Goexit()
	}()


	for i := 0; i < 5; i++ {
		fmt.Println("im main")
		time.Sleep(time.Second)
	}
}
