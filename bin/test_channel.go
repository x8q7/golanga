package main

import (
	"fmt"
	"time"
)

func send(ch chan interface{}) {
	//ch <- 10
	return
}

func receive(ch chan interface{}) {
	s := <-ch
	fmt.Println(s)
	return
}

func main() {
	ch1 := make(chan interface{})

	go receive(ch1)
	go send(ch1)
	time.Sleep(1 * time.Second)
	fmt.Println("close start")
	close(ch1)
	fmt.Println("close end")
	time.Sleep(5 * time.Second)
}
