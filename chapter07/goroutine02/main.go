package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//互斥锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	for i := 1; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock() //go build -race main.go
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
