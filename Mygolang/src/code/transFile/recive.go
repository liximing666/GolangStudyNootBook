package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Recive struct {

}

func (this *Recive) Start() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("等待客户端连接......")

	for {
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go Handel(connect)
	}



}

func Handel(connect net.Conn) {
	defer connect.Close()
	fmt.Println("客户端连接成功......")

	for  {
		//读取客户端传过来的文件名
		buf := make([]byte, 4096)
		n, err := connect.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0 {
			fmt.Println("sendor leave")
			return
		}

		fileName := string(buf[:n])
		fmt.Println(fileName)

		//给客户端发送回执
		connect.Write([]byte("has revive filename"))

		//接收客户端传过来的文件内容
		ReciveFile(connect, fileName)


	}

}

func ReciveFile(connect net.Conn, fileNamee string) {
	//按照文件名创建新文件
	file, err := os.Create(fileNamee)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	//从网络中读取文件
	for {
		buf := make([]byte, 10240)
		n, err := connect.Read(buf)
		if n == 0 {
			fmt.Println("end")
			return
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("end")
			}else {
				fmt.Println(err)
			}
			return
		}

		fileContent := buf[:n]
		file.Write(fileContent)

		fmt.Println("文件写入完毕......")
	}
}

func NewRecive() *Recive {
	return &Recive{}
}
