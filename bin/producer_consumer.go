package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 90; i += 10 {
		ch <- i
	}
	close(ch) // 发送完成后关闭通道
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch { // 自动在通道关闭后退出
		fmt.Println(value)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
