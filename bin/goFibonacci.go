package main

import "fmt"

//
//从示例 6.13 fibonacci.go 的斐波那契程序开始，制定解决方案，使斐波那契周期计算独立到协程中，并可以把结果发送给通道。
//
//结束的时候关闭通道。main() 函数读取通道并打印结果：goFibonacci.go

func fibonacci(ch chan int, n int) {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	input := make(chan int)
	go fibonacci(input, 10)

	for i := range input {
		fmt.Println(i)
	}
}
