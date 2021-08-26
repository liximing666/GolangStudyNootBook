package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//按字符串写
	f, _ := os.OpenFile("C:\\Users\\ximin\\Desktop\\字节写.txt", os.O_RDWR, 7)
	f.WriteString("ha123")

	//按位置写 先seek 后writeAt
	ret, _ := f.Seek(5, io.SeekStart)
	fmt.Println(ret)
	f.WriteAt([]byte("this is write"), ret)
	//按字节写



	//***********************************************

	//按行读
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
	//按字节读

	img, _ := os.ReadFile("D:\\壁纸")
	fmt.Println("du wan le")
	f1, _ := os.OpenFile("C:\\Users\\ximin\\Desktop\\img.docx", os.O_RDWR, 7)
	f1.Write([]byte(img))
	fmt.Println("xie wan le")
	f1.Close()

	defer f.Close()
}
