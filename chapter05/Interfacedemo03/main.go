package main

import "fmt"

type Usb interface {
	Say()
}
type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say()")
}
func main() {
	var stu Stu = Stu{}
	u := &stu
	u.Say()

	//错误 interface默认是指针类型
	//var u1 Usb
	//u1.Say()

}
