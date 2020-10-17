package main

import "fmt"

func getSum(n1, n2 int) int {
	return n1 + n2
}

//函数也是一种数据类型
func myFunc(funcvar func(int, int) int, num1, num2 int) int {
	return funcvar(num1, num2)
}

type myFuncType func(int, int) int

func myFunc2(funcvar myFuncType, num1, num2 int) int {
	return funcvar(num1, num2)
}

func main() {

	a := getSum
	fmt.Printf("a的类型\n%T\ngetSum的类型\n%T\n", a, getSum)

	res := a(10, 40)
	fmt.Println(res)

	res2 := myFunc(getSum, 50, 60)
	fmt.Println(res2)

	//自定义数据类型
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 20
	num1 = myInt(num2)
	fmt.Println("num1=", num1, "num2=", num2)

	res3 := myFunc2(getSum, 1, 2)
	fmt.Println(res3)
}
