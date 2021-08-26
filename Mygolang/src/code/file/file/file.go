package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:\\Users\\ximin\\Desktop\\t.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("hellofile\n"))
	f.WriteString("im ok")


	f1, err := os.OpenFile("C:\\Users\\ximin\\Desktop\\hahaha.txt", os.O_RDWR, 7)
	if err != nil {
		fmt.Println(err)
	}
	f1.WriteString("this hahah\n")


	defer f.Close()

}
