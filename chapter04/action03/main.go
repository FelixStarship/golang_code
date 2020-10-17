package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (s Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", s.Name, s.Age, s.Score)
}

func (s Student) SetScore(score float64) {
	s.Score = score
}

type Pupil struct {
	//TODO   Student Student   嵌套结构体加了名称 访问时就必须加上名称  组合
	Student //匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试..........")
}

type Graduate struct {
	Student Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中.........")
}

func main() {

	//TODO   Student Student   嵌套结构体加了名称 访问时就必须加上名称 (组合)
	//p:=&Pupil{}
	//p.Student.Name="tom~"
	//p.Student.Age=18
	//p.Student.Score=89.9
	//p.testing()
	//p.Student.SetScore(98.9)
	//p.Student.ShowInfo()

	var p Pupil
	p.Age = 19
	p.Name = "tom~"
	p.Age = 90
	p.Score = 99.99
	p.SetScore(90.00)
	p.ShowInfo()

}
