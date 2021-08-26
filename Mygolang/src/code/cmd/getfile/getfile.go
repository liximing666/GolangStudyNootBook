package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("请命令行输入文件的路径")
	lists := os.Args
	if len(lists) != 2 {
		fmt.Println("输入不合法")
		return
	}
	path := lists[1]

	//按照路径获得文件信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())

}
