package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// 初始化统计变量
	var (
		charCount   int                         // 字符总数（含空格，不含\r\n）
		wordCount   int                         // 单词总数
		lineCount   int                         // 行数
		inWord      bool                        // 标记是否处于单词中
		stopSignal  = false                     // 标记是否输入了 'S' 结束符
		inputReader = bufio.NewReader(os.Stdin) // 手动读取输入的 Reader
	)

	fmt.Println("请输入内容（输入单个字符 'S' 并回车结束输入）：")

	// 循环读取每行输入（直到输入 'S' 或出错）
	for {
		// ReadString('\n')：读取到换行符为止，返回的字符串包含 '\n'
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入出错：%v\n", err)
			return
		}
		lineCount++ // 每读取一行，行数+1

		// 遍历当前行的每个字符（包含 '\n'，需手动处理）
		for _, c := range input {
			// 1. 检查结束符：仅当行内容是 "S\n" 时触发（单个 'S' + 换行）
			if c == 'S' && len(input) == 2 && input[1] == '\n' {
				stopSignal = true
				break // 退出字符遍历
			}

			// 2. 统计字符数：排除 \r 和 \n
			if c != '\r' && c != '\n' {
				charCount++
			}

			// 3. 统计单词数：空白字符（空格/制表符等）分隔
			if unicode.IsSpace(c) {
				// 跳过 '\n' 和 '\r'（避免误判单词结束）
				if c == '\n' || c == '\r' {
					continue
				}
				inWord = false
			} else if !inWord {
				// 非空白字符 + 不在单词中 → 新单词
				inWord = true
				wordCount++
			}
		}

		// 检测到结束符，退出循环
		if stopSignal {
			break
		}
	}

	// 输出统计结果（最后一行的 'S' 不计入）
	fmt.Println("\n===== 统计结果 =====")
	fmt.Printf("1. 字符个数（含空格，不含\\r/\\n）：%d\n", charCount)
	fmt.Printf("2. 单词个数（空白字符分隔）：%d\n", wordCount)
	fmt.Printf("3. 行数：%d\n", lineCount-1) // 减去最后一行的 'S' 行
}
