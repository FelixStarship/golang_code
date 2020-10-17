package main

import "fmt"

func test(arr *[3]int) {
	(*arr)[1] = 90
	fmt.Printf("arr=%p    arr[0]=%p\n", &arr, &arr[0])
}

func test02(n *int) {
	*n = 10
}

func main() {

	var intArr [3]int
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Printf("intArr数组的地址=%p  intArr[0]数组的地址=%v\n", &intArr, &intArr[0])

	//初始化数组的方式
	var numArr [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr", numArr)

	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02", numArr02)

	var numArr03 = [...]int{1, 2, 3}

	fmt.Println("numArr03", numArr03)

	var numArr04 = [...]int{1: 100, 0: 233, 3: 211313}
	fmt.Println("numArr04", numArr04)

	for i := 0; i < len(numArr04); i++ {
		fmt.Println(numArr04[i])
	}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	var str [3]string = [3]string{"松江", "无用", "成语"}
	for i, v := range str {
		fmt.Println(i, v)
	}

	var arr = [3]int{1, 2, 3}
	test(&arr)
	fmt.Printf("arr=%p  arr[0]=%p", &arr, &arr[0])

}
