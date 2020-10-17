package main

import "fmt"

var (
	func1 = func(n1, n2 int) int {
		return n1 + n2
	}(1, 2)
)

func main() {

	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)

	//函数变量
	a := func(n1, n2 int) int {
		return n1 + n2
	}

	res2 := a(1, 2)
	fmt.Printf("res2=%v  type=%T\n", res2, a)

	fmt.Printf("func1=%v  type=%T", func1, func1)
}
