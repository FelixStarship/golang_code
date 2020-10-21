package main

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world!")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() err", err)
		}
	}()
	var m map[int]string
	m[0] = "hello world!"
}
func main() {

	go SayHello()
	go test()

	time.Sleep(time.Second * 5)
}
