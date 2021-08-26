package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
}

func (this *Server) Start()  {

	//指定服务器
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))//listen本质是socket
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("等待客户端")



	connect, _ := listen.Accept() // 本质上是服务器的通信socket，监听客户端连接
	fmt.Println("连接成功")

	//被动接受数据
	buf := make([]byte, 4096)//读过来的数据存到buf
	n, err := connect.Read(buf)

	fmt.Println(string(buf[:n]))

	//给客户端发送回执
	connect.Write([]byte("is feed back"))

	defer listen.Close()
	defer connect.Close()

}

func NewServer(ip string, port int) *Server {
	 return &Server{Ip: ip, Port: port}
}
