package main

import (
	"fmt"
	"math"
)

// Polar 极坐标
type Polar struct {
	Radius float64
	Angle  float64 // 角度（度）
}

// Cartesian 笛卡尔坐标
type Cartesian struct {
	X float64
	Y float64
}

// 协程：从 channel1 读极坐标，转成笛卡尔坐标发送到 channel2
func convert(pChan <-chan Polar, cChan chan<- Cartesian) {
	for p := range pChan {
		// 角度转弧度
		rad := p.Angle * math.Pi / 180
		x := p.Radius * math.Cos(rad)
		y := p.Radius * math.Sin(rad)
		cChan <- Cartesian{X: x, Y: y}
	}
}

func main() {
	fmt.Println("=== 极坐标 → 笛卡尔坐标 转换器 ===")

	// 两个通道
	channel1 := make(chan Polar)
	channel2 := make(chan Cartesian)

	// 启动转换协程
	go convert(channel1, channel2)

	// 交互输入
	var r, angle float64
	fmt.Print("请输入半径: ")
	fmt.Scan(&r)
	fmt.Print("请输入角度(度): ")
	fmt.Scan(&angle)

	// 发送到 channel1
	channel1 <- Polar{Radius: r, Angle: angle}

	// 从 channel2 接收结果
	cart := <-channel2

	fmt.Printf("\n转换结果：\n")
	fmt.Printf("极坐标(%.2f, %.2f°) → 笛卡尔坐标(%.4f, %.4f)\n",
		r, angle, cart.X, cart.Y)

	close(channel1)
	close(channel2)
}
