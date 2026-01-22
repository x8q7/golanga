package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 运算符自定义类型（修正拼写：isVaild → isValid）
type operator string

const (
	add      operator = "+"
	subtract operator = "-"
	multiply operator = "*"
	divide   operator = "/"
)

// 修正拼写：isVaild → isValid，优化校验逻辑
func (op operator) isValid() bool {
	switch op {
	case add, subtract, multiply, divide:
		return true
	default:
		return false
	}
}

// 计算器结构体（新增 reset 方法用于重置步骤）
type calculator struct {
	number1 int
	number2 int
	op      operator
	step    int // 0:待输入num1, 1:待输入num2, 2:待输入op, 3:计算完成
}

// 初始化计算器
func NewCal() *calculator {
	return &calculator{step: 0} // 其他字段默认0即可
}

// 重置计算器状态（输入错误或计算完成后重置）
func (c *calculator) reset() {
	c.number1 = 0
	c.number2 = 0
	c.op = ""
	c.step = 0
}

// 输入处理（核心修复：类型转换替代断言、完善逻辑）
func (c *calculator) Input(input string) (isFinal bool, ok bool) {
	// 优先处理退出指令
	if input == "q" {
		fmt.Println("bye bye~")
		os.Exit(0)
	}

	fmt.Printf("当前输入：%s，步骤：%d\n", input, c.step)
	ok = true
	isFinal = false

	switch c.step {
	case 0: // 输入第一个数字
		num1, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("错误：输入的不是合法整数 → %v\n", err)
			ok = false
			c.reset() // 重置步骤，重新输入
			return
		}
		// 校验数值范围（最大值 999999）
		if num1 < 0 || num1 > 999999 {
			fmt.Println("错误：数字超出范围（0~999999）")
			ok = false
			c.reset()
			return
		}
		c.number1 = num1
		c.step = 1 // 进入下一步
	case 1: // 输入第二个数字
		num2, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("错误：输入的不是合法整数 → %v\n", err)
			ok = false
			c.reset()
			return
		}
		if num2 < 0 || num2 > 999999 {
			fmt.Println("错误：数字超出范围（0~999999）")
			ok = false
			c.reset()
			return
		}
		c.number2 = num2
		c.step = 2 // 进入下一步
	case 2: // 输入运算符
		// 核心修复：string → operator 显式类型转换（而非断言）
		op := operator(input)
		if !op.isValid() {
			fmt.Println("错误：无效运算符（仅支持 +、-、*、/）")
			ok = false
			c.reset()
			return
		}
		c.op = op
		isFinal = true // 输入完成，触发计算
		c.step = 3     // 标记计算完成
	default: // 计算完成后重置，重新开始
		c.reset()
		ok = false
		return
	}
	return
}

// 计算逻辑（完善错误处理）
func (c *calculator) Calc() (int, error) {
	switch c.op {
	case add:
		return c.number1 + c.number2, nil
	case subtract:
		return c.number1 - c.number2, nil
	case multiply:
		return c.number1 * c.number2, nil
	case divide:
		if c.number2 == 0 {
			return 0, errors.New("错误：除数不能为 0")
		}
		return c.number1 / c.number2, nil
	default:
		return 0, errors.New("错误：无效的运算符")
	}
}

func main() {
	calc := NewCal()
	inputReader := bufio.NewReader(os.Stdin) // 只创建一次，避免重复初始化
	fmt.Println("===== 逆波兰式计算器 =====")
	fmt.Println("输入规则：依次输入 数字1 → 数字2 → 运算符（+、-、*、/），输入 q 退出")
	fmt.Println("数值范围：0 ~ 999999\n")

	for {
		fmt.Print("请输入：")
		// 读取输入并清洗（去除 \n、\r、空格）
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入失败：%v\n", err)
			return
		}
		input = strings.TrimSpace(input) // 清洗输入（关键：先清洗再判断 q）

		// 空输入跳过
		if input == "" {
			continue
		}

		// 处理输入
		isFinal, ok := calc.Input(input)
		if !ok {
			fmt.Println("请重新输入！\n")
			continue
		}

		// 输入完成，执行计算
		if isFinal {
			result, err := calc.Calc()
			if err != nil {
				fmt.Printf("计算失败：%v\n", err)
			} else {
				fmt.Printf("\n✅ 计算结果：%d %s %d = %d\n\n", calc.number1, calc.op, calc.number2, result)
			}
			calc.reset() // 重置计算器，准备下一次计算
		}
	}
}
