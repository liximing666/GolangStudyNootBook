package main

type Hero struct {
	Name string
	Age int


}
func NewHero(name string, age int) *Hero {
	return &Hero{Name: name, Age: age}
}



type HeroList struct {

	Heros []*Hero

}

func NewHeroList() *HeroList {
	return &HeroList{}
}


func (heroList *HeroList) Add(hero *Hero)  {
	heroList.Heros = append(heroList.Heros, hero)
}



func (heroList *HeroList) Len() int {
	return len(heroList.Heros)
}

//定制排序
func (heroList *HeroList) Less(i, j int) bool {
	return  heroList.Heros[i].Age < heroList.Heros[j].Age
}

func (heroList *HeroList) Swap(i, j int) {
	tmp := heroList.Heros[i]
	heroList.Heros[i] = heroList.Heros[j]
	heroList.Heros[j] = tmp
}




