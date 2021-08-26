package main

import "fmt"

//空接口是一种万能的数据类型，类似于object类型
func Myfunc(arg interface{})  {
	fmt.Println("this func is called")
	fmt.Println(arg)

	//判断空接口的具体数据类型，用断言
	_, ok := arg.(int)
	fmt.Println(ok)
}

func main() {
	Myfunc(66)

}
