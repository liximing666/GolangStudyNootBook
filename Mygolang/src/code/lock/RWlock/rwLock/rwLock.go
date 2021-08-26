package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


/*
1 读锁是共享的 写锁是独占的 写锁的优先级高（r w 一起来先w）
2 不要吧 互斥锁 读写锁 管道 混用 可能会隐形死锁
3 锁只有一把（拿到读锁 写锁再来也会被阻塞）

多 gorutine 通信 推荐读写锁

读写锁 和 channel保护共享数据的区别
	读写锁 一个写 多个读 1 : n (效率高)
	channel 一个写一个读 1 : 1

*/

func Reader(i int)  {
	time.Sleep(time.Second)

	lock.RLock()
	fmt.Printf("------%d读到了%d\n", i, values)
	lock.RUnlock()
}

func Writer(i int)  {
	random := rand.Intn(1000)

	lock.Lock()
	values = random
	fmt.Printf("%d写入了%d\n", i, random)
	lock.Unlock()
}

var values int
var lock sync.RWMutex


func main() {

	for i := 1; i <= 5; i++ {
		go Reader(i)
	}

	for i := 1; i <= 5; i++ {
		go Writer(i)
	}

	for  {
		;
	}

}
