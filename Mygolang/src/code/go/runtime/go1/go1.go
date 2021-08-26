package main

import (
	"fmt"
	"runtime"
)

func main() {
	//设置参与运算的cpu核心数
	runtime.GOMAXPROCS(4)

	for  {
		fmt.Print(1)

		go func() {
			fmt.Print(0)
		}()

	}


}
