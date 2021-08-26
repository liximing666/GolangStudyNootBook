package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	ch := make(chan int)

	go func() {
		timer.Reset(1 * time.Second)//重置
		fmt.Println(<- timer.C)
		 timer.Stop()//停止
		 ch <- 666
	}()

	for {
		<- ch
		break
	}

}
