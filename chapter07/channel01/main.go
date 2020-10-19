package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 90
	//intChan<-899   deadlock!  不会自动扩容不能超过其容量
	fmt.Printf("channel len=%v  cap=%v\n", len(intChan), cap(intChan))
	var intNum int
	//从管道取值
	intNum = <-intChan
	<-intChan
	<-intChan
	//<-intChan    deadlock!
	fmt.Println("intNum=", intNum)
	fmt.Printf("channel len=%v  cap=%v\n", len(intChan), cap(intChan))
}
