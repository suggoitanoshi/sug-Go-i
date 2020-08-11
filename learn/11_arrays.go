package main

import "fmt"

func main() {
	var a [2]string // deklarasi array, [n]T dengan n jumlah dan T tipe data
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // deklarasi dengan isi
	fmt.Println(primes)

	var s []int = primes[1:4] // cara slicing array, mirip di python: [low (inklusif):high (eksklusif)]
	fmt.Println(s)
	// perlu diingat hasil slicing tetap menunjuk ke array yang sama
	s[0] = 1
	fmt.Println(primes)

	// slice literal: untuk inisialisasi array, tapi tanpa length
	t := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(t)
	// untuk slice, dapat menghapus value yang tidak mau diset
	// default low: 0, high: slice length

	// konsep length dan capacity
	// length: jumlah elemen di slice nya
	// capacity: jumlah elemen di array aslinya
	u := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(len(u), cap(u), u)
	u = u[:0]
	fmt.Println(len(u), cap(u), u)
	// untuk zero slice, value-nya nil, length dan capacitynya 0
	var v []int
	fmt.Println(v == nil)

	// untuk membuat dynamic length array, pakai make()
	// cara pakai: make([]T, len, cap)
	w := make([]int, 5)
	fmt.Println(cap(w), len(w), w)

	// untuk menambahkan elemen ke akhir, gunakan append()
	// cara pakai: append(slice, vs)
	var x []int
	fmt.Println(x)
	x = append(x, 10)
	fmt.Println(x)
	x = append(x, 20, 30, 40)
	fmt.Println(x)

	// untuk mengiterasi elemen array, pakai range
	// range menerima slice, mengembalikan index dan kopi elemen
	var pow = []int{1, 2, 4, 8, 16}
	for index, value := range pow {
		fmt.Println(index, value)
	}
}
