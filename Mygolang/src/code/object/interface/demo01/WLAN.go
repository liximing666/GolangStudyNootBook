package main

type Wlan interface {
	Start()
	Stop()
	Show(interface{}) interface{}
}
