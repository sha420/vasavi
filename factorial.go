package main

import "fmt"

func factorial(n int) int {
	if n != 0 {
		return factorial(n-1) * n
	} else {
		return 1
	}
}

func main() {
	fmt.Println(factorial(10))
}
