package main

import "fmt"

func integers(ch chan int) {
	a, b := 0, 1
	ch <- a
	go func() {
		for {
			ch <- b
			a, b = b, a+b
		}
	}()
}

func generatE(ch chan int) int {
	return <-ch
}

func main() {

	ch := make(chan int)
	go integers(ch)
	for i := 0; i < 10; i++ {
		res := generatE(ch)
		fmt.Println(res)
	}

}
