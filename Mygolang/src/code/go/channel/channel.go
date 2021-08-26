package main

import "fmt"

//channel用来保证进程之间的通信，自带同步
func main() {

	//定义一个无缓冲区channel
	c1 := make(chan int)

	go func () {
		defer fmt.Println("协程执行完毕")
		fmt.Println("协程正在执行")
		c1 <- 666
	}()
	x :=  <-c1
	fmt.Println(x)
	fmt.Println("main 结束")




}
