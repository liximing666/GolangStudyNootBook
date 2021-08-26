package main

//结构体标签是对结构体属性的的手动说明，可以通过反射机制把标签拿出来
import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `info:"this is a name" doc:"this is a doc"`
	age int `info:"this is a age"`
}

func StruceTag(s interface{})  {
	Mytype := reflect.TypeOf(s)
	//通过反射把标签取出来
	for i := 0; i < Mytype.NumField(); i++ {
		info := Mytype.Field(i).Tag.Get("info")
		doc := Mytype.Field(i).Tag.Get("doc")
		fmt.Printf("info : %s , doc : %s\n", info, doc)
	}




}

func main() {
	s := Person{name: "kity", age: 16}
	StruceTag(s)
}
