package main

import "fmt"

type Vertex struct {
	x, y int
}

var m map[string]Vertex

func main() {
	// map: memetakan key ke value
	// semacam associative array
	m = make(map[string]Vertex)
	m["origin"] = Vertex{0, 0}
	fmt.Println(m["origin"])

	// map literal harus menggunakan key dan value
	n := map[string]Vertex{
		"origin": Vertex{0, 0},
		"2pow2":  Vertex{2, 4},
	}
	fmt.Println(n)
	// untuk mutasi map
	n["4pow2"] = Vertex{4, 16} // assign key ke elemen baru
	e := n["origin"]           // mengambil elemen
	delete(n, "2pow2")         // menghapus elemen
	_, ok := n["3pow2"]        // mengambil dan mengetes elemen, apakah ada di array
	fmt.Println(e)
	fmt.Println(ok)
	fmt.Println(n)
}
