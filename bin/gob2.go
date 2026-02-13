// gob2.go
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	file, _ := os.OpenFile("vcard.gob", os.O_RDONLY, 0777)
	defer file.Close()

	dec := gob.NewDecoder(file)

	var decodedVCard VCard
	err := dec.Decode(&decodedVCard)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	// 4. 打印解码结果
	fmt.Println("\n✅ 解码后的 VCard 数据：")
	fmt.Printf("姓名：%s %s\n", decodedVCard.FirstName, decodedVCard.LastName)
	fmt.Println("地址列表：")
	for i, addr := range decodedVCard.Addresses {
		fmt.Printf("  %d. 类型：%s，城市：%s，国家：%s\n", i+1, addr.Type, addr.City, addr.Country)
	}
	fmt.Printf("备注：%s\n", decodedVCard.Remark)
}
