package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	var p1 Person
	p1.Name = "小明"
	p1.Age = 20

	var p2 *Person = &p1
	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	(*p2).Name = "tom~"
	fmt.Printf("p2.Name=%v p1.Nmae=%v \n", p2.Name, p1.Name)
	fmt.Printf("p2.Name=%v p1.Nmae=%v \n", (*p2).Name, p1.Name)

	fmt.Printf("p1的地址：%p \n", &p1)
	fmt.Printf("p2的地址：%p  p2的值：%p \n", &p2, p2)
}
