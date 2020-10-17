package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float32
	ptr    *int
	slice  []int
	map1   map[int]int
}

func main() {

	ptr1 := 10
	var p1 Person = Person{ptr: &ptr1}
	fmt.Printf("ptr1=%v\n", &ptr1)
	*p1.ptr = 90
	fmt.Printf("p1.ptr=%v  p1.ptr=%v  p1=%v\n", p1.ptr, &p1.ptr, p1)

	var p2 *Person = new(Person)
	(*p2).Name = "test"
	p2.Name = "more"
	fmt.Println(p2)

	var p3 *Person = &Person{} //结构体指针
	p3.Name = "name"           //go底层做了取值运算(*p3)
	//*取值拿到person
	(*p3).Name = "age"

}
