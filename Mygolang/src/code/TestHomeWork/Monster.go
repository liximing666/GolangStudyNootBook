package TestHomeWork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {

	Name string
	Age int
	Skill string

}

func (m *Monster) Store()  {
	//序列化对象
	json_m, _ := json.Marshal(*m)
	//存到文件
	f, _ := os.OpenFile("C:/Users/ximin/Desktop/sei.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	f.Write(json_m)
	fmt.Println("store ok")
	defer f.Close()
}

func (m *Monster) ReStore() *Monster {
	f, _ := os.OpenFile("C:/Users/ximin/Desktop/sei.txt", os.O_RDONLY,0600)
	defer f.Close()

	bytes, _ :=ioutil.ReadAll(f)

	//反序列化
	m1 := Monster{}
	json.Unmarshal(bytes, &m1)

	return &m1

}
