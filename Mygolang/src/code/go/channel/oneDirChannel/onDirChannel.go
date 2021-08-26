package main

import "fmt"

// 管道定义赋值拿出来的是 地址，结构体不是

func Send(s chan <- int)  {
	s <- 666
	close(s)
}

func Recive(r <- chan int)  {
	num := <- r
	fmt.Println(num)
}




func main() {
	C := make(chan int)

	go func() {
		Send(C)
	}()

	Recive(C)

	var a chan <- int = C
	fmt.Println(C,a)


}
