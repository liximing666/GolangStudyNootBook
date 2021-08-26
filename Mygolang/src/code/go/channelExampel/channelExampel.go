package main

import "fmt"

func feb1(n int) int {

	if n <= 2 {
		return 1
	}

	febLIst := make([]int,n)
	febLIst[0] = 1
	febLIst[1] = 1


	for i := 2; i < len(febLIst); i++ {
			febLIst[i] = febLIst[i-1] + febLIst[i-2]
		}

	return febLIst[len(febLIst)-1]
}

func feb(c ,quit chan int)  {
	x, y := 1, 1

	//用for select 监管多个通道
	for {
		select {

		case c <- x:
			tmp := x + y
			x = y
			y = tmp

		case <- quit:
			fmt.Println("quit")
			return
		}


	}
}

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	feb(c, quit)


}
