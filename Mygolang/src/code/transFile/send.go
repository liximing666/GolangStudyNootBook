package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Sendor struct {

}

func Send(connect net.Conn ) {
	addr := connect.RemoteAddr()
	fmt.Printf("%s, plz input file path......\n", addr)


	lists := os.Args
	if len(lists) != 2 {
		fmt.Println("file path error")
	}
	filePath := lists[1]

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileInfo.Sys()
}

func (this *Sendor) Start() {
	//获取文件名
	fmt.Println("plz input file abstract path......")
	lists := os.Args
	if len(lists) != 2 {
		fmt.Println("file path error")
		return
	}
	filePath := lists[1]

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileName := fileInfo.Name()

	//拨号服务器
	connect, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("send err", err)
	}

	//把文件名发送给服务器
	connect.Write([]byte(fileName))

	//读服务器的反馈
	buf := make([]byte, 4096)
	n, err := connect.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	if n == 0 {
		fmt.Println("sendor leave")
		return
	}
	feedBack := buf[:n-1]
	if feedBack != nil {
		// 发送文件
		sendFile(connect, filePath)
	}



	defer connect.Close()


}

func sendFile(connect net.Conn, path string) {
	defer connect.Close()
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	for  {
		//读取文件内容
		buf := make([]byte, 10240)
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送完毕......")
			}else {
				fmt.Println(err)
			}
			return
		}

		content := buf[:n-1]

		//发送文件
		connect.Write(content)
	}

}

func NewSendor() *Sendor {
	return &Sendor{}
}