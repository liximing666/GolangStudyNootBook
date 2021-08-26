package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)


// 把结构体转成json
type Movie struct {

	//打jsonTge 在解析json的时候用的Tag做json的key

	Title string `json:"title"`
	Price int `json:"price"`
	Year int `json:"year"`
	Actors []string `json:"actors"`
}

func main() {

	//把结构体变成json     Marshal()
	movie := Movie{Title: "inception", Price: 40, Year: 2014, Actors: []string{"li", "jan"}}
	JsonList, _ := json.Marshal(movie)
	fmt.Printf("JsonList = %s", JsonList)

	//把json变结构体    Unmarshal()
	movie1 := &Movie{}
	json.Unmarshal(JsonList, movie1)

	Mymovie := reflect.ValueOf(*movie1)
	for i := 0; i < Mymovie.NumField(); i++ {
		fmt.Println(Mymovie.Field(i))
	}
	fmt.Printf("%v",*movie1 )
}
