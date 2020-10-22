package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	_ "k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type Stop struct {
}

func main() {
	stopCh := make(chan struct{})
	i := 0
	go wait.Until(func() {
		fmt.Printf("----%d----\n", i)
		i++
	}, time.Second, stopCh)
	<-stopCh
}
