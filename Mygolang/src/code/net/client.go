package main

import (
	"fmt"
	"net"
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
	connet.Write([]byte("hi server"))

	//接收服务器发送回来的回执
	buf1 := make([]byte, 4096)//读过来的数据存到buf
	n, _ := connet.Read(buf1)
	fmt.Println(string(buf1[:n]))


}

func NewClient(ip string, port int) *Client {
	return &Client{Ip: ip, Port: port}
}
