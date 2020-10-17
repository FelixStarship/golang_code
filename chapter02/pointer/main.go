package main

import "fmt"

func main() {

	//&取址  *取值
	var i = 10
	fmt.Printf("i的地址：%v\n", &i)

	//指针变量
	var ptr *int = &i
	*ptr = 20
	fmt.Printf("ptr:=%v    type:%T\n", ptr, ptr)
	fmt.Printf("ptr的地址:%v\n", &ptr)
	//先取地址 然后在*取值
	fmt.Printf("ptr的值：%v\n", *ptr)
	fmt.Printf("i的值：%v\n", i)

	//var a int=90
	////指针变量接收的是地址不是值
	//var test *int=a
	////类型不匹配
	//var ptr *float32=float32(&a)

	//案例演示
	var a int = 300
	var b int = 400
	//指向a
	var ptr1 *int = &a
	*ptr1 = 100
	//重新指向b
	ptr1 = &b
	*ptr1 = 200
	fmt.Printf("a=%d,b=%d,*ptr1=%d", a, b, *ptr1)

}
