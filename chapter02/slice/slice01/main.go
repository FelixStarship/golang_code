package main

import "fmt"

func main() {

	var intArr = [...]int{1, 2, 3, 4, 56}
	fmt.Printf("intArr=%T  intArr=%p  intArr=%p  intArr=%v   intArr=%v  intArr=%v\n", intArr, &intArr, &intArr[1], len(intArr), cap(intArr), intArr)
	//切片是数组的一个引用
	slice := intArr[1:3]
	slice[0] = 90
	//slice[1]=800
	//slice[2]=899
	fmt.Printf("slice=%T  slice=%p  slice=%v  slice=%v  slice=%p  slice=%v\n", slice, &slice, len(slice), cap(slice), &slice[0], slice)
	fmt.Printf("intArr=%v  slice=%v", intArr[1], slice[0])

	//数组越界 切片
	var slice1 []int = make([]int, 1)
	slice1[0] = 10
	slice1[1] = 30
	slice1[2] = 90
	fmt.Printf("slice1=%v", slice1)

}
