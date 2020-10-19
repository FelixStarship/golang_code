package main

import "fmt"

func main() {

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//intChan<-300    管道关闭后不能在写入数据
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//管道遍历
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	//deadlock!  在遍历管道时如果channel没有关闭会死锁
	//close(intChan2)

	for v := range intChan2 {
		fmt.Println("v=", v)
	}

	//for  {
	//	v,ok:=<-intChan2
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}
}
