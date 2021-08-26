package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := make([]int, 3, 6)
	fmt.Printf("len = %d  cap = %d %v\n", len(slice1), cap(slice1), slice1)



	slice1 = append(slice1, 6)
	fmt.Printf("len = %d  cap = %d %v\n", len(slice1), cap(slice1), slice1)

	for i := 0; i < 3; i++ {
		slice1 = append(slice1, i+1)
	}

	fmt.Printf("len = %d  cap = %d %v\n", len(slice1), cap(slice1), slice1)



	sort.Ints(slice1)
	fmt.Printf("len = %d  cap = %d %v\n", len(slice1), cap(slice1), slice1)

	var a = slice1[0:len(slice1)-1]
	fmt.Printf("slice1 = %v, a = %v\n", slice1, a)



	//截取后指向的是同一个对象，一改都改，是浅拷贝
	a[0] = 6
	fmt.Printf("slice1 = %v, a = %v\n", slice1, a)

	//copy方法进行深拷贝
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 999
	fmt.Printf("a = %v b = %v\n", a, b)




}
