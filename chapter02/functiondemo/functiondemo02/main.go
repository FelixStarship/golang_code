package main

import "fmt"

//指针变量
func test(n1 *int) {
	fmt.Printf("n1=%v\n", &n1)
	*n1 = *n1 + 1
	fmt.Printf("n1=%T\n", *n1)
	fmt.Println("test() n1=", *n1)
}

func main() {
	num := 10
	fmt.Printf("num的地址：%v\n", &num)
	test(&num)
	fmt.Println("main() num=", num)
}
