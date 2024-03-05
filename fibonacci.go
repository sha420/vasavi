package main

import (
	"fmt"
)

func fibo(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	n := 10
	fmt.Printf("The %d-th Fibonacci number is: %d\n", n, fibo(n))
}
