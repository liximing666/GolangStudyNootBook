package main

import (
	"fmt"
	"net"
	"runtime"
)

func Handel(connect net.Conn, i int)  {
	addr := connect.RemoteAddr()
	fmt.Printf("shoudao%d, exit to break\n", i, addr)
	defer connect.Close()

	for {
		buf := make([]byte, 4096)//buf 的最后一个元素是个换行符
		n, _ := connect.Read(buf)

		if n == 0 { // 读到0 说明读写的channel已经关闭了
			fmt.Println("come exit")
			return
		}
		content := string(buf[:n-1])//buf 的最后一个元素是个换行符

		if content == "exit" {
			fmt.Println("exit")
			runtime.Goexit()
			return
		}else {
			fmt.Println(content + "form", i, " input exit to break")
			//向客户端发送反馈
			connect.Write([]byte("is feedback\n"))
		}
	}


}

type Server struct {
	Ip string
	Port int
}

func (this *Server) Start()  {
	i := 1
	//指定服务器
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))//listen本质是socket
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("等待客户端")

	for {
		connect, _ := listen.Accept() // 本质上是服务器的通信socket，监听客户端连接
		fmt.Println("连接成功")

		go Handel(connect, i)

		i ++
	}



	defer listen.Close()


}

func NewServer(ip string, port int) *Server {
	return &Server{Ip: ip, Port: port}
}