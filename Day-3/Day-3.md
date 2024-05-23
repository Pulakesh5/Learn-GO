# Day -3

- Implementing sqrt() using newton’s method:
    - ```go
        func Sqrt(x float64) float64 {
        	const lim = 1e-6
        	sqrtValue := 1.0
        	prev := sqrtValue
        	for {
        		prev = sqrtValue
        		sqrtValue -= (sqrtValue*sqrtValue - x) / (2 * sqrtValue)
        		if sqrtValue-prev <= lim {
        			break
        		}
        		fmt.Println(sqrtValue)
        	}
        	return sqrtValue
        }
        ```
        
- GO’s switch case format is similar to C, C++ etc. But one thing that sets it apart is it only runs the selected part, not the cases that follow.
    - ```go
        func main() {
        	switch os := runtime.GOOS; os {
        	case "darwin":
        		fmt.Println("Mac OS X")
        	case "linux":
        		fmt.Println("Linux.")
        	default:
        		// freebsd, openbsd,
        		// plan9, windows...
        		fmt.Printf("%s.\n", os)
        	}
        }
        ```
        
- switch a condition is same as `switch true` . Clean way to write long `if-else` statements.
    -   ```go
        func main() {
        t := time.Now()
        switch {
        case t.Hour() < 12:
        fmt.Println("Good morning!")
        case t.Hour() < 17:
        fmt.Println("Good afternoon.")
        default:
        fmt.Println("Good evening.")
        }
        ```
        
- A defer statement defers the execution of a function until the surrounding function returns. Deferred function calls are pushed onto a stack. When surrounding function returns, deferred calls are executed in LIFO order
    - ```go
        package main
        
        import "fmt"
        
        func main() {
        	defer fmt.Println("world")
        
        	fmt.Println("hello")
        }
        ```
        
        ```go
        package main
        
        import "fmt"
        
        func main() {
        	fmt.Println("counting")
        
        	for i := 0; i < 10; i++ {
        		defer fmt.Println(i)
        	}
        
        	fmt.Println("done")
        }
        
        ```
        
        Ouput 
        
        ```
        counting
        done
        9
        8
        7
        6
        5
        4
        3
        2
        1
        0
        ```
        

---

## More Types - structs, slice and maps

- Pointers in GO are just like they are in C++, but without pointer arithmetic. Zero value of pointer in GO is `nil`
    - ```go
        func main() {
        	i, j := 42, 2701
        
        	p := &i         // point to i
        	fmt.Println(*p) // read i through the pointer
        	*p = 21         // set i through the pointer
        	fmt.Println(i)  // see the new value of i
        
        	p = &j         // point to j
        	*p = *p / 37   // divide j through the pointer
        	fmt.Println(j) // see the new value of j
        }
        // Output: 
        // 42
        // 21 
        // 73
        ```
        
- A struct is like a structure in C++. To access fields of a struct `.` notation is used. It allows us to use pointer as a variable without dereferencing [in terms of notation]
    - ```go
        package main
        
        import "fmt"
        
        type Vertex struct {
        	X int
        	Y int
        }
        
        func main() {
        
        	v := Vertex{1, 2}
        	fmt.Println(Vertex{1, 2})
        	p := &v
        	p.X = 1e9
        	fmt.Println(v)
        }
        // Output:
        // {1 2}
        // {1000000000 2}
        ```
        
- Example of struct literal
    - ```go
        var (
        	v1 = Vertex{1, 2}  // has type Vertex
        	v2 = Vertex{Y: 1}  // Y:0 is implicit
        	v3 = Vertex{}      // X:0 and Y:0
        	p  = &Vertex{1, 2} // has type *Vertex
        )
        
        func main() {
        	fmt.Println(v1, p, v2, v3)
        }
        // Output: {1 2} &{1 2} {0 1} {0 0}
        ```
        
- The expression `var a [10]int` declares `a` as an array of ten integers.
    - ```go
        func main() {
        	var a [2]string
        	a[0] = "Hello"
        	a[1] = "World"
        	fmt.Println(a[0], a[1])
        	fmt.Println(a)
        
        	primes := [6]int{2, 3, 5, 7, 11, 13}
        	fmt.Println(primes)
        }
        // Output:
        // Hello World
        // [Hello World]
        // [2 3 5 7 11 13]
        ```
        
    - Slice is like vector in C++, dynamic and flexible. The type `[]T` is a slice with elements of type `T` . `a[low : high]` This selects the `a[low]` but excludes `a[high]` .
    - ```go
            func main() {
            	primes := [6]int{2, 3, 5, 7, 11, 13}
            
            	var s []int = primes[1:4]
            	fmt.Println(s)
            }
            // Output: [3 5 7]
            ```
