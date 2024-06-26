package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{Y: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {

	// v := Vertex{1, 2}
	// fmt.Println(Vertex{1, 2})
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
