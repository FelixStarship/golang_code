package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作了")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作了")
}

func (p Phone) End() {
	fmt.Println("test end")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作了")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作了")
}

type Computer struct {
}

func (c Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}

func main() {

	c := Computer{}
	p := Phone{}
	ca := Camera{}
	var u Usb = p
	fmt.Println(u)
	c.Working(p)
	c.Working(ca)

}
