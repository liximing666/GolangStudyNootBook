package main

import "fmt"

func swap(pa *int, pb *int)  {
	var tmp = 0
	tmp = *pa
	*pa = *pb
	*pb = tmp
}

func swap1(a int, b int) (int, int)  {
	var tmp = 0
	tmp = a
	a = b
	b = tmp
	return a, b
}

func main() {
	a := 100
	add := &a
	add1 := &add
	fmt.Println(add)
	fmt.Println(*add)
	fmt.Println(**add1)
	fmt.Println("****************************************")

	b := new(int)
	fmt.Println(*b)

	c := 1999
	d := 9991
	fmt.Println(c, d)
	swap(&c, &d)
	fmt.Println(c, d)

	e := 1999
	f := 9991
	fmt.Println(e, f)
	e, f = swap1(e, f)
	fmt.Println(e, f)

}
