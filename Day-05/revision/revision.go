package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s",
// 		e.When, e.What)
// }

// func run() error {
// 	myErr := MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// 	return &myErr
// 	// return &MyError{
// 	// 	time.Now(),
// 	// 	"it didn't work",
// 	// }
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

/*
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
*/

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%s is (%d years old)", p.Name, p.Age)
// }

// func main() {
// 	p := Person{"Alice", 30}
// 	fmt.Println(p) // Output: Alice (30 years old)
// 	fmt.Println(p.String())
// 	// above can be thought of as fmt.Println(p.String())
// }

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	var s []int = primes[:]
// 	fmt.Println(s, len(s), cap(s))
// 	s = s[1:4]
// 	fmt.Println(s, len(s), cap(s))
// }

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {

// 	v := Vertex{1, 2}
// 	fmt.Println(Vertex{1, 2})
// 	p := &v
// 	// p.X = 1e9 // NO need for (*p).X
// 	(*p).Y = 1e9
// 	fmt.Println(v)
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(split(17))
// }

// output: 7 10

// func pow(x, n, lim float64) float64 {
// 	fmt.Println("pow called with", x, n, lim)
// 	if v := math.Pow(x, n); v < lim {
// 		fmt.Printf("%f is less than %f\n", v, lim)
// 		return v
// 	} else {
// 		fmt.Printf("%f is greater than %f\n", v, lim)
// 		// fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	// can't use v here, though
// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }
