package main

import "fmt"

// tel 协程：循环发送数字到通道
func tel(ch chan int) {
	// 循环发送 1~10
	for i := 1; i <= 10; i++ {
		ch <- i // 发送数字
	}
	close(ch) // 发完关闭通道，告诉 main 没有更多数据
}

func main() {
	// 创建无缓冲通道
	ch := make(chan int)

	// 启动协程
	go tel(ch)

	// main 从通道接收并打印，直到通道关闭
	for v := range ch {
		fmt.Println(v)
	}
}
