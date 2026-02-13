package main

import (
	"fmt"
	"log"
	"time"
)

// 模拟请求
type Request struct {
	ID int
}

// 模拟处理请求
func handleReq(r Request) {
	log.Printf("处理请求：%d\n", r.ID)
}

func main() {
	// 1. 带缓冲通道：扛暴增流量（队列容量 100）
	queue := make(chan Request, 100)

	// 2. 每秒处理 10 个请求 = 每 100ms 处理一个
	const rate = 10
	interval := time.Second / time.Duration(rate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 3. 消费者：定时从队列取任务执行
	go func() {
		for range ticker.C {
			fmt.Printf("queue length: %d \n", len(queue))
			select {
			case req := <-queue:
				handleReq(req)
			default:
				// 队列空，不阻塞
			}
		}
	}()

	// 4. 模拟瞬间暴增的请求（生产者）
	for i := 0; i < 50; i++ {
		queue <- Request{ID: i}
	}
	log.Println("50 个突发请求已全部入队")

	// 等待处理完
	time.Sleep(6 * time.Second)
}
