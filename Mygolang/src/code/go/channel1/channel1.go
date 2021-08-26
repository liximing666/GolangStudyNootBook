package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个有缓冲区channel，缓冲的大小不能变，溢出会阻塞
	c2 := make(chan int, 3)
	fmt.Println(len(c2), cap(c2))

	go func() {
		fmt.Println("正在执行协程......")

		for i := 0; i < 5; i++ {
			c2 <- i
			fmt.Println("写入", i)
		}
		fmt.Println("协程完毕......")

		//关闭channel
		close(c2)
	}()

	time.Sleep(3*time.Second)

	for i := 0; i < 5; i++ {
		num := <- c2
		fmt.Println("输出", num)
	}
	fmt.Println("main完毕")
	fmt.Println(len(c2), cap(c2))

}