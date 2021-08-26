package main

import "fmt"

func show(map1 map[int]string)  {

	for k , v := range map1 {
		fmt.Printf("k = %d  v = %s\n", k, v)
	}
}

func copyMap(map1 map[int]string) map[int]string {
	map2 := map[int]string{}

	for k , v := range map1 {
		map2[k] = v
	}
	return map2
}

func main() {
	map1 := map[int]string{
		1 : "java",
		2 : "c++",
		3 : "python",
		4 : "php",
	}

	show(map1)

	delete(map1, 4)

	show(map1)

	fmt.Println(map1[3])

	map1[4] = "golang"

	show(map1)

	//深拷贝map没有现成方法要自己手动实现
	map2 := copyMap(map1)

	fmt.Printf("map1 = %v , map2 = %v\n", map1, map2)




}
