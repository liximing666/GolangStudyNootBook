package main

func main() {

	apple := Server{Name: "APPLE"}

	c := Computer{}
	apple.Use(&c)
	apple.Use(&Phone{})



}
