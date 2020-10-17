package main

import "fmt"

type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
	Say()
}
type Student struct {
}
type Monster struct {
}

//隐式实现interface
func (s Student) Say() {
	fmt.Println("Student Say()")
}
func (m Monster) Hello() {
	fmt.Println("Monster Hello~")
}
func (s Monster) Say() {
	fmt.Println("Monster Say~")
}

//只要是自定义类型都可以实现接口
type integer int

func (i integer) Say() {
	fmt.Println("integer Say()")
}
func main() {
	var stu Student
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	//自定义类型实现多个接口
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
