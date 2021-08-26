package main

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Ip string
	Port int
}

func (this *Client) Start() {
	//客户端拨号去连接服务器，返回客户端通信的套接字
	connet, err := net.Dial("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connet.Close()


	//客户端给服务器写数据
	go SendToServer(connet)
	//客户端接收服务器的反馈
	go ReciveFeadBack(connet)
}

func ReciveFeadBack(connect net.Conn) {

	for {

		buf := make([]byte, 4096)
		n, err := connect.Read(buf)
		if n == 0 {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}else {
			fmt.Println(string(buf[:n]))
		}

	}
}

func NewClient(ip string, port int) *Client {
	return &Client{Ip: ip, Port: port}
}

func SendToServer(connect net.Conn )  {
	fmt.Println("pls input content")

	for {
		//从键盘读取内容
		buf := make([]byte,4096)
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			fmt.Println("input err")
			continue
		}


		content := string(buf[:n-1])

		if content == "exit" {
			return
		}else {
			_, err := connect.Write([]byte(content))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}