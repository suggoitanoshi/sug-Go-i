package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// ini closure:
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // masing-masing memegang "variabel" sendiri

	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
}
