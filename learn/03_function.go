package main

import "fmt"

// deklarasi: nama dulu baru tipe
// func add(x int, y int) int {
// atau lebih singkatnya, karena tipenya sama:
func add(x, y int) int {
	// keyword return untuk mengembalikan nilai ke pemanggil
	return x + y
}

// fungsi bisa mengembalikan lebih dari satu nilai
func swap(a, b string) (string, string) {
	return b, a
}

// fungsi bisa mengembalikan nilai bernama
func named(a int) (b, c int) {
	b = a - 10
	c = a + 10
	// "Naked" return statement
	// Mengembalikan nilai bernama dari deklarasi fungsi
	return
}

func main() {
	fmt.Println(add(69, 420))           // 489
	fmt.Println(swap("Hello", "World")) // World Hello
	fmt.Println(named(365))             // 355 375
}
