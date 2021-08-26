package main

import (
	"fmt"
	"sync"
	"time"
)

//标准库实现的互斥锁 可以自己用channel来实现，对共享区域加互斥锁

var lock sync.Mutex
func Printer(s string)  {
	for _, char := range s {
		fmt.Printf("%c",char)
		time.Sleep(300 * time.Microsecond)
	}
}

func Person1()  {
	lock.Lock()
	Printer("hello")
	lock.Unlock()

}
func Person2(q chan bool)  {
	time.Sleep(1 * time.Second)
	lock.Lock()
	Printer("world")
	lock.Unlock()
	q <- true
}



func main() {

	quit := make(chan bool)

	go Person1()

	go Person2(quit)

	<- quit
}
