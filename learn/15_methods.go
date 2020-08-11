package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method untuk tipe, didahului "receiver" argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs)
}
