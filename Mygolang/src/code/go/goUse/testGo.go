package main

import (
	"fmt"
	"runtime"
	"time"
)

func T1() {
	for i := 0; i < 100; i++ {
		fmt.Println("im T1 , from  ", i)
		time.Sleep(1*time.Second)//表达时间的方式
	}
	runtime.Goexit()
}

func T2() {
	for i := 0; i < 100; i++ {
		fmt.Println("im T2 , from  %d ", i)
		time.Sleep(1*time.Second)
	}
	runtime.Goexit()
}

func main() {

	//用 go 关键字来开线程
	go T1()
	go T2()

	for i := 0; i < 100; i++ {
		fmt.Println("im main , from  %d ", i)
		time.Sleep(1*time.Second)

	}


}
