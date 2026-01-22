package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 核心：定义存储单行数据的结构体（标题/价格/数量）
type Product struct {
	Title  string // 标题
	Price  string // 价格（暂存字符串，可按需转float64）
	Amount string // 数量（暂存字符串，可按需转int）
}

func main() {
	// 1. 打开文件
	filePath := "products.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败：%v\n", err)
		return
	}
	defer file.Close() // 确保文件最终关闭

	// 2. 初始化结构体切片，存储所有产品数据
	var products []Product

	// 3. 使用 bufio.Scanner 按行读取文件（正确的逐行读取方式）
	scanner := bufio.NewScanner(file)
	lineNum := 0 // 记录行号，方便定位错误
	for scanner.Scan() {
		lineNum++
		// 获取当前行的完整内容（自动去除换行符）
		line := scanner.Text()
		// 跳过空行
		if strings.TrimSpace(line) == "" {
			continue
		}

		// 4. 按分号分割字段
		fields := strings.Split(line, ";")
		// 校验字段数量（必须有3个：标题/价格/数量）
		if len(fields) != 3 {
			fmt.Printf("第 %d 行格式错误：字段数=%d，期望=3，内容=%s\n", lineNum, len(fields), line)
			continue
		}

		// 5. 清洗字段（去除首尾空格，避免多余空白）
		title := strings.TrimSpace(fields[0])
		price := strings.TrimSpace(fields[1])
		amount := strings.TrimSpace(fields[2])

		// 6. 创建Product结构体并加入切片
		product := Product{
			Title:  title,
			Price:  price,
			Amount: amount,
		}
		products = append(products, product)
	}

	// 7. 检查扫描过程中是否出错（非EOF错误）
	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件出错：%v\n", err)
		return
	}

	// 8. 打印所有产品数据（格式化输出）
	fmt.Println("===== 产品数据列表 =====")
	if len(products) == 0 {
		fmt.Println("文件中无有效产品数据")
		return
	}
	// 打印表头
	fmt.Printf("%-20s %-10s %-5s\n", "标题", "价格", "数量")
	fmt.Println("----------------------------------------")
	// 打印每个产品
	for i, p := range products {
		fmt.Printf("%d. %-18s %-10s %-5s\n", i+1, p.Title, p.Price, p.Amount)
	}
}
