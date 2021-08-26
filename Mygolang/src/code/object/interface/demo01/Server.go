package main

import "fmt"

type Server struct {
	Name string
}

func (server *Server) Use(wlan Wlan)  {
	wlan.Start()
	wlan.Stop()
	fmt.Println(wlan.Show(666))
}
