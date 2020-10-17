package main

import (
	"fmt"
	"go_code/chapter04/action02"
)

type Person struct {
	Name string
}

type Dog struct {

}

func (p * Person)test()  {
	(*p).Name="jack"
	fmt.Println(*p)
}

func main() {
   var p Person
   p.Name="tom~"
   (&p).test()
   fmt.Println(p)

   stu:=model.NewStudent("tom~",88.9)
   fmt.Println(*stu)


}
