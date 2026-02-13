package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// randomBit 生成一个随机位：0 或 1
func randomBit() int {
	// 生成 [0, 2) 范围内的安全随机数
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}

func main() {
	fmt.Println("无限随机位生成器（0/1），按 Ctrl+C 停止：")
	// 无限循环生成位
	for {
		fmt.Print(randomBit())
	}
}
