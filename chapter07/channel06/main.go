package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v, ok := <-intChan:
			fmt.Printf("从intChan读取的数据%v ok=%v \n", v, ok)
		case v, ok := <-stringChan:
			fmt.Printf("从stringChan读取的数据%v ok=%v\n", v, ok)
			//default:
			//	fmt.Println("读取不到数据了")
			//	return
		}
	}
}
