package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//实现对对象的定制排序 实现Interface接口放进sort
func main() {
	heroList := NewHeroList()

	for i := 0; i < 10; i++ {
		hero := NewHero("hero" + string(rand.Intn(100)), rand.Intn(100))
		heroList.Add(hero)
	}

	for i := 0; i < heroList.Len(); i++ {
		fmt.Println(heroList.Heros[i])

	}

	fmt.Println("sort*********************")

	sort.Sort(heroList)
	for i := 0; i < heroList.Len(); i++ {
		fmt.Println(heroList.Heros[i])

	}

}
