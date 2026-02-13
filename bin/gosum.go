package main

import "fmt"

func sum(x int, y int, ch chan int) {
	s := x + y
	ch <- s
}

func main() {
	a := 10
	b := 20
	ch := make(chan int)
	go sum(a, b, ch)
	res := <-ch
	fmt.Println(res)
}
