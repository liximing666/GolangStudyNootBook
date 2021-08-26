package main

import (
	"fmt"
	"time"
)

//timer 定时到了，操作系统写入时间 然后自己读取
func main() {
	fmt.Println(time.Now())
	t := time.NewTimer(6 * time.Second)
	fmt.Println(<- t.C)

	fmt.Println(<- time.After(3 * time.Second))
}
