package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {

	//结构体的所有字段在内存中是连续的
	r1 := Rect{Point{x: 1, y: 2}, Point{x: 3, y: 4}}
	fmt.Printf("r1.leftUp.x 地址=%p  r1.leftUp.y 地址=%p   r1.rightDown.x 地址=%p  r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	fmt.Printf("leftUp=%p   rightDown=%p\n", &r1.leftUp, &r1.rightDown)
	r2 := Rect2{leftUp: &Point{
		x: 10,
		y: 20,
	}, rightDown: &Point{
		x: 30,
		y: 40,
	}}
	//指针指向的地址不一定连续
	fmt.Printf("leftUp=%p   rightDown=%p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("leftUp  指向地址=%p   rightDown  指向地址=%p\n", r2.leftUp, r2.rightDown)
}
