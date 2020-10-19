package main

import (
	"fmt"
	"runtime"
)

func main() {

	numCpu := runtime.NumCPU()

	fmt.Println("numCpu", numCpu)
	//设置可同时执行的最大CPU数
	runtime.GOMAXPROCS(numCpu - 1)
}
