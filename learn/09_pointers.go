package main

import "fmt"

func main() {
	i, j := 69, 420
	var p *int      // pointer ke integer, nilainya nil
	p = &i          // ambil pointer ke i
	fmt.Println(*p) // baca isi dari variabel yang ditunjuk p, yaitu i
	*p = 21         // isi variabel yang ditunjuk p, yaitu  i, dengan 21
	fmt.Println(i)  //baca isi variabel i secara langsung

	// contoh lebih banyak
	p = &j
	*p = *p / 10
	fmt.Println(j)
}
