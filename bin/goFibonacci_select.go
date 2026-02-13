package main

import (
	"fmt"
	"time"
)

//
//从示例 6.13 fibonacci.go 的斐波那契程序开始，制定解决方案，使斐波那契周期计算独立到协程中，并可以把结果发送给通道。
//
//结束的时候关闭通道。main() 函数读取通道并打印结果：goFibonacci.go
//使用 select 语句来写，并让通道退出 (gofibonacci_select.go)

func fibonacci(ch chan int, quit chan interface{}, n int) {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-quit:
			break
		default:
			fmt.Println("default")
		}
	}
	close(ch)
}

func main() {
	input := make(chan int)
	quit := make(chan interface{})
	go fibonacci(input, quit, 10000)

	go func() {
		for i := range input {
			fmt.Println(i)
		}
	}()
	time.Sleep(200 * time.Microsecond)
	close(quit)
}
