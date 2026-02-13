package main

import (
	"fmt"
	"strconv"
)

//写一个 ConvertInt64ToInt() 函数把 int64 值转换为 int 值，如果发生错误（提示：参见 4.5.2.1 节）就 panic() 。然后在函数 IntFromInt64 中调用这个函数并 recover()，返回一个整数和一个错误。请测试这个函数！

func ConvertInt64ToInt(i interface{}) (r int) {
	switch nr := i.(type) {
	case int:
		r = int(nr)
	default:
		panic("input canot transfer to int")
	}
	return
}

func IntFromInt64(input interface{}) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = fmt.Errorf("")
		}
	}()
	result = ConvertInt64ToInt(input)
	return
}

func main() {
	// 正常情况
	v1 := int64(999)
	i1, e1 := IntFromInt64(strconv.Itoa(int(v1)))
	fmt.Printf("v1=%d => int=%d, err=%v\n", v1, i1, e1)

}
