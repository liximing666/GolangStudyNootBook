package main

import (
	"reflect"
	"fmt"
)

type phoneInf interface {
	call()
}

type Phone struct {
	name string
	price int
}

func (this *Phone) Call() {
	fmt.Println("im call")
}

func main() {
	//通过反射机制，把value和type拿出来（pair） ,拿出来的value 和 type 都是对象

	value1 := reflect.ValueOf(Phone{name: "iphone", price: 6999})
	fmt.Println(value1)
	type1 := reflect.TypeOf(Phone{name: "iphone", price: 6999})
	fmt.Println(type1)
	t := reflect.TypeOf(type1)
	fmt.Println(t)

	//遍历
	for i := 0; i < value1.NumField(); i++ {
		fmt.Println(value1.Field(i))
	}
	//通过反射拿到值
	n1 := value1.Field(0)
	n := value1.FieldByName("name")
	fmt.Println(n, n1)

	for i := 0; i < type1.NumMethod(); i++ {
		m := type1.Method(i)
		fmt.Println(m.Name, m.Type)
	}


}
