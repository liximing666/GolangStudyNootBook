package main

import (
	"fmt"
	"time"
)

func main() {
/* 自己实现的定时器
timer := time.NewTimer(1 * time.Second)
	go func() {
		for {
			t, ok := <- timer.C
			fmt.Println(t)
			if ok {
				timer.Reset(1 * time.Second)
			}else {
				fmt.Println("error")
				break
			}
		}
	}()
*/
	// Ticker 是用来循环周期计时的
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println(time.Now())
	ch := make(chan int)


	go func() {
		i := 0
		for i < 6 {
			fmt.Println(<- ticker.C)
			i++
		}
		ch <- 666
	}()

	for {
		<- ch
		break
	}

}
