// panic_defer.go
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g. ") // 1
	g(0)
	fmt.Println("Returned normally from g.") //  6
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")   // 3
		panic(fmt.Sprintf("%v", i)) // 5
	}
	defer fmt.Println("Defer in g", i) // 4
	fmt.Println("Printing in g", i)    // 0,1,2,
	g(i + 1)
}
