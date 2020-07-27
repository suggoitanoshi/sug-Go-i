package main

import "fmt"

func main() {
	defer fmt.Println("World") // Ditunda eksekusinya sampai fungsi luarnya selesai

	// untuk defer yang banyak, yang pertama dieksekusi adalah yang terakhir didefer
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Hello")
}
