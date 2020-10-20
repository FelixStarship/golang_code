package main

import (
	"sync"
)

func main() {

	/*workers:=3
		ch:=make(chan struct{})
		worker:= func() {
			ch<- struct{}{}
		}
		leader:= func() {
			cnt:=0
			for range ch{
	          cnt++
				if cnt==workers {
	               break
				}
			}
			close(ch)
		}
		go leader()
		for i:=0;i<workers;i++ {
			go worker()
		}*/

	wg := sync.WaitGroup{}
	workers := 3
	wg.Add(workers)
	worker := func() {
		defer wg.Done()
	}
	leader := func() {
		wg.Wait()
	}
	go leader()
	for i := 0; i < workers; i++ {
		go worker()
	}

}
