package main

import "fmt"

func main() {

	//go func() {
	//	fmt.Println("你好，世界！")
	//}()
	//
	//select {}

	c := make(chan int) //无缓冲管道
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("写入", i)
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("读取", num)
	}

}
