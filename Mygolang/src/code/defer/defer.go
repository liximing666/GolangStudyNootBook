package main

import "fmt"

func main() {
	//defer 关键字定义最后一步的操作，defer多个从上到下进栈
	//return先 defer后，defer是在函数生命周期结束之后
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("step 1")
	fmt.Println("step 2")
}
