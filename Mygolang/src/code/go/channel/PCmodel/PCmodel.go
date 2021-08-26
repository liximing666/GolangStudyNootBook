package main

import (
	"fmt"
	"time"
)

type Productor struct {

}
func (p *Productor) Send(v int, ch chan <- int) {
	ch <- v
	fmt.Println("生产者生产了----------", v)
}
func NewProductor() *Productor {
	return &Productor{}
}

type Consumer struct {

}

func (c *Consumer) Recive(ch  <- chan int) {
	num, ok := <- ch
	if ok {
		fmt.Println("消费者消费了-------------", num)
	}else {
		fmt.Println("error")
	}
}
func NewConsumer() *Consumer {
	return &Consumer{}
}

func main() {
	signal := make(chan int)
	ch := make(chan int, 6)
	productor := NewProductor()
	comsumer := NewConsumer()

	go func() {
		for i := 1; i <= 100; i++ {
			productor.Send(i, ch)
			time.Sleep(3 * time.Microsecond)
		}
		signal <- 666

	}()

	go func() {
		for {
			comsumer.Recive(ch)
			time.Sleep(1 * time.Second)
		}

	}()

	for  {
		<- signal
		break
	}

}
