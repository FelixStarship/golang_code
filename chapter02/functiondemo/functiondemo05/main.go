package main

import (
	"fmt"
	"strconv"
	"strings"
)

//闭包
func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Println(str)
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("test"))
	fmt.Println(f2("test.peg"))
}
