package main

import (
	"fmt"
	"runtime"
	"time"
)

//超时机制 管道阻塞一段时间就超时
func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
				case num := <- ch:
					fmt.Println(num)
				case <- time.After(3 * time.Second): // 能读出来说明已经超时了
					fmt.Println("time out")
					quit <- true
					goto label //跳转语句 函数内生效
			}
		}
		label: runtime.Goexit()
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}

	<- quit

}
