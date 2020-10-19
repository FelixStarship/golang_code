package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allCat := make(chan interface{}, 3)
	allCat <- 3
	allCat <- "tom~"
	cat := Cat{
		Name: "小花猫",
		Age:  20,
	}
	allCat <- cat
	<-allCat
	<-allCat
	newCat := <-allCat
	fmt.Printf("newCat=%T   newCat=%v\n", newCat, newCat)
	//类型断言
	a := newCat.(Cat)
	fmt.Printf("a=%T   a.Name=%v\n", a, a.Name)
}
