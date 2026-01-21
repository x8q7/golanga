package main

import (
	"fmt"
	"strconv"
)

//在练习 10.16 和 10.17 中我们开发了一些栈结构类型。但是它们被限制为某种固定的内建类型。现在用一个元素类型是 interface{}（空接口）的切片开发一个通用的栈类型。
//
//实现下面的栈方法：
//
//Len() int
//IsEmpty() bool
//Push(x interface{})
//Pop() (interface{}, error)
//Pop() 改变栈并返回最顶部的元素；Top() 只返回最顶部元素。
//
//在主程序中构建一个充满不同类型元素的栈，然后弹出并打印所有元素的值。
//

type Stack struct {
	elements []interface{}
}

func (this *Stack) String() string {
	str := "["
	for i, _ := range this.elements {
		element := this.elements[i]
		if i != 0 {
			str += ", "
		}
		switch t := element.(type) {
		case int:
			str += strconv.Itoa(t)
		case string:
			str += t
		case bool:
			str += strconv.FormatBool(t)
		}
	}
	str += "]"
	return str
}

func (this *Stack) Len() int {
	return len(this.elements)
}

func (this *Stack) IsEmpty() bool {
	return len(this.elements) <= 0
}

func (this *Stack) Push(elements ...interface{}) {
	this.elements = append(this.elements, elements...)
}

func (this *Stack) Pop() (result interface{}, ok bool) {
	if this.IsEmpty() {
		return nil, false
	}
	result = this.elements[len(this.elements)-1:][0]
	this.elements = this.elements[:len(this.elements)-1]
	return result, true
}

func main() {
	stack := &Stack{elements: make([]interface{}, 0)}
	stack.Push(123)
	fmt.Println(stack)
	stack.Push("456")
	stack.Push(false)
	fmt.Println(stack)
	last, ok := stack.Pop()
	if !ok {
		fmt.Println("stack no member")
	}
	fmt.Println(last)
	fmt.Println(stack)

}
