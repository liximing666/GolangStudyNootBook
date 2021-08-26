package main

import (
	"fmt"
	"os"
)

func main() {
	//键盘获得输入

	buf := make([]byte, 4096)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n-1]))


	fmt.Println("******************************")


	//命令行获得参数,可执行文件的名字是第一个

	list := os.Args
	fmt.Println(list[1])
}
