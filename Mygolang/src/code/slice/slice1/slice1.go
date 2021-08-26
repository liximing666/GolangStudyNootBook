package main

import "fmt"

func show(array []int)  {
	for _, v := range array{
		println(v)
	}
}

func main() {
//普通数组是静态的， slice是动态的数组
	
	//普通数组
	var array [10]int
	array[0] = 10

	arr1 := [10]int {1,2,3,4}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
//range 就是enumerate方法
	for index, value := range arr1{
		fmt.Println(index, value)
	}



	//动态数组 slice 长度不固定
	var slice1 = []int {6,6,6,6}
	show(slice1)

}
