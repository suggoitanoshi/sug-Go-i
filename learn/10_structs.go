package main

import "fmt"

// struct adalah kumpulan "field"
type Vertex struct {
	X int // ini field
	Y int // ini field juga
}

func main() {
	v := Vertex{1, 2} // inisialisasi struct
	v.X = 20          // akses menggunakan dot
	p := &v
	p.Y = 19            // Bisa akses tanpa harus dereference
	v2 := Vertex{X: 69} // Y:0 implisit, bila ingin default semua, gunakan tanpa inisialisasi
	fmt.Println(v, v2)
}
