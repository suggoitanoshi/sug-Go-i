package main

import "fmt"

func main() {
	// mirip seperti bahasa modern
	// namun tidak perlu kurung
	x := 6
	if true {
		fmt.Println("Hello")
		x = 4
	}
	if false {
		fmt.Println("World")
		x = 8
	}
	// if menerima init juga, seperti for
	if i := 20; x < 5 {
		fmt.Println(x, i) // 4 20
	} else { // else harus berada di baris yang sama dengan kurung kurawal penutup
		// deklarasi init di if juga dapat diakses di else
		fmt.Println(i, x)
	}
}
