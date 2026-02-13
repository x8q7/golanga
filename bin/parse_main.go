package main

import (
	"fmt"
	"text/template/parse"
)

func main() {
	// 测试场景1：正常解析
	input1 := "1 2 3 4 5"
	nums1, err1 := parse.Parse(input1)
	fmt.Printf("输入：%q → 结果：%v，错误：%v\n", input1, nums1, err1)

	// 测试场景2：包含非法字符（解析失败）
	input2 := "1 a 3 4"
	nums2, err2 := parse.Parse(input2)
	fmt.Printf("输入：%q → 结果：%v，错误：%v\n", input2, nums2, err2)
	// 断言为自定义 ParseError，获取详细信息
	if pe, ok := err2.(*parse.ParseError); ok {
		fmt.Printf("  错误详情：索引=%d，错误字符串=%q，原始错误=%v\n", pe.Index, pe.Word, pe.Err)
	}

	// 测试场景3：空输入（无内容可解析）
	input3 := ""
	nums3, err3 := parse.Parse(input3)
	fmt.Printf("输入：%q → 结果：%v，错误：%v\n", input3, nums3, err3)

	// 测试场景4：多个空格分隔
	input4 := "  10  20  30  "
	nums4, err4 := parse.Parse(input4)
	fmt.Printf("输入：%q → 结果：%v，错误：%v\n", input4, nums4, err4)
}
