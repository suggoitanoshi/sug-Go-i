package main

import "fmt"

func main() {
	sum := 0
	// strukturnya mirip bahasa modern
	// init; condition; post
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 0
	// yang harus ada hanya condition, init dan post boleh tidak ada
	for sum < 100 { // atau for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
}
