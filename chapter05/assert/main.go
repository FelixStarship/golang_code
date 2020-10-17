package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	//var a interface{}
	//var point Point
	//a=point
	////类型断言
	//var b Point
	//b=a.(Point)
	//fmt.Println(b)
	var x interface{}
	var b2 float32
	x = b2
	if a, ok := x.(float32); ok {
		fmt.Println(a)
	} else {
		fmt.Println("convert fail")
	}

}
