package main

import "fmt"

func operate(fn func(int, int) int) int {
	return fn(4, 9)
}

func main() {
	mult := func(a, b int) int {
		return a * b
	}
	fmt.Println(operate(mult))
}
