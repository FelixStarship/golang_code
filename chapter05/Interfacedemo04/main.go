package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

//切片
type HeroSlice []Hero

func (h HeroSlice) Len() int {
	return len(h)
}
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age > h[j].Age
}
func (h HeroSlice) Swap(i, j int) {
	//temp:=h[i]
	//h[i]=h[j]
	//h[j]=temp
	h[i], h[j] = h[j], h[i]
}

func main() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		h := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, h)
	}
	fmt.Println("-----------排序前-----------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("-----------排序后-----------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
