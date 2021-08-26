package main

import "fmt"

func main() {
	var map1 map[string]string = make(map[string]string)
	map1["k1"] = "v1"
	map1["k2"] = "v2"
	map1["k3"] = "v3"

	fmt.Printf("map1 = %v\n", map1)

	map2 := make(map[int]int,3)
	map2[1] = 6
	fmt.Printf("map2 = %v\n", map2)

	map3 := map[int]string{
		1 : "java",
		2 : "c++",
		3 : "python",
	}
	fmt.Printf("map3 = %v\n", map3)



}
