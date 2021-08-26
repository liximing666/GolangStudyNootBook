package main

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		for {
			select {
			case n1 := <- ch1:
				ch2 <- n1
			case n2 := <- ch2:
				ch1 <- n2
			}
		}
	}()

	<- ch1
	<- ch2


}
