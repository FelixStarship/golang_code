package main

import "fmt"

type Monkey struct {
	Name string
}
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}
type LittleMonkey struct {
	Monkey
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生来会爬树")
}
func (l *LittleMonkey) Flying() {
	fmt.Println(l.Name, "通过学习，会飞翔...........")
}
func (l *LittleMonkey) Swimming() {
	fmt.Println(l.Name, "通过学习，会游泳.........")
}
func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	(&monkey).climbing()
	monkey.Flying()
	monkey.Swimming()

	//属性继承  行为接口  实现接口可以看作是对继承的一种补充
	var b BirdAble = &monkey
	b.Flying()

	//is-a  like-a   接口在于设计 关系更加松散 解耦
}
