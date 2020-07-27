package main

import "fmt"

var a, b, c bool              // Deklarasi variabel diawali var dan dapat diakhiri dengan tipe
var d, e int = 1, 2           // d = 1, e = 2
var f, g, h = 69, true, "Yo!" // f = 69, g = true, h = "Yo!"
// tipe data:
// bool, string,
// int, int8, int16, int32, int64
// uint, uint16, uint32, uint64, uintptr
// rune -> unicode code point, int32
// float32 float64
// complex64 complex128
// contoh untuk complex number
var z = 2 + 3i

// untuk konversi v ke tipe T, gunakan T(v)
var y = float32(d)

// untuk konstanta, gunakan keyword const, tidak bisa dideklarasi dengan :=
const euler = 3.14

func main() {
	i := 420                                     // tipe implisit, hanya dapat digunakan dalam fungsi
	fmt.Println(a, b, c, d, e, f, g, h, i, y, z) // false false false 1 2 69 true Yo!
}
