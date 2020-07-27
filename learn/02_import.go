package main

// untuk banyak import bisa difaktorkan
// separatornya newline
import (
	"fmt"
	"math"
)

func main() {
	// nama yang bisa diakses Pi, karena yang diexport itu yang diawali huruf kapital
	fmt.Println("Pi: ", math.Pi)
}
