package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}
type Computer struct {
}

func (c Computer) Working(u Usb) {
	//类型断言
	if v, ok := u.(Phone); ok {
		v.Call()
	}
	u.Start()
	u.Stop()
}
func (p Phone) Start() {
	fmt.Println("手机开始工作了")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作了")
}
func (p Phone) Call() {
	fmt.Println("手机在打电话...........")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作了")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作了")
}
func main() {
	//多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{Name: "vivo"}
	usbArr[1] = Phone{Name: "iphone"}
	usbArr[2] = Camera{Name: "索尼"}
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
