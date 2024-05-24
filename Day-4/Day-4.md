# Day-4

---

## Methods

- Go doest not support classes. To, define methods that works on a certain type, we use methods [`= function + receiver`].
  `receiver` argument is a instance of the type the method works on.
  - ```go
     type Vertex struct {
     	X, Y float64
     }

     func (v Vertex) Abs() float64 {
     	ans := math.Sqrt(v.X*v.X + v.Y*v.Y)
     	v.X = v.X-2
     	v.Y = v.Y-2
     	return ans;
     }

     func main() {
     	v := Vertex{3.0, 4.0}
     	fmt.Println(v.Abs())
     	fmt.Println(v)
     }

     /*
     5
     {3 4} // no change in the original variable by the method
     */
    ```
- You can declare methods with pointer receivers. With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behaviour as for any other function argument.) See how `Abs` method works on a copy of the original `Vertex` value, while `Scale` method works on the value to which the receiver points.
  - ```go
    func (v Vertex) Abs() float64 {
    	ans := math.Sqrt(v.X*v.X + v.Y*v.Y)
    	v.X = v.X-2
    	v.Y = v.Y-2
    	return ans;
    }

    func (v *Vertex) Scale(f float64) {
    	v.X = v.X * f
    	v.Y = v.Y * f
    }

    func main() {
    	v := Vertex{3, 4}
    	fmt.Println(v)
    	v.Scale(10)
    	fmt.Println(v)
    	fmt.Println(v.Abs())
    	fmt.Println(v)
    }

    /*
    {3 4}
    {30 40}
    50
    {30 40}
    */
    ```
- Functions with pointer argument must take a pointer, while methods with pointer receivers take either a value or a pointer as the receiver when they are called.
  - ```go
    func (v *Vertex) Scale(f float64) {
    	v.X = v.X * f
    	v.Y = v.Y * f
    }

    func ScaleFunc(v *Vertex, f float64) {
    	v.X = v.X * f
    	v.Y = v.Y * f
    }

    func main() {
    	v := Vertex{3, 4}
    	v.Scale(2)
    	ScaleFunc(&v, 10)

    	p := &Vertex{4, 3}
    	p.Scale(3)
    	ScaleFunc(p, 8)

    	fmt.Println(v, p)
    }
    ```
  Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale`  method has a pointer receiver.
- Similar thing happens with functions/ methods with value argument.
  - ```go
    var v Vertex
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // Compile error!
    ```
    ```go
    var v Vertex
    fmt.Println(v.Abs()) // OK
    p := &v
    fmt.Println(p.Abs()) // OK
    ```
    ```go
    func (v Vertex) Abs() float64 {
    	return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func AbsFunc(v Vertex) float64 {
    	return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func main() {
    	v := Vertex{3, 4}
    	fmt.Println(v.Abs())
    	fmt.Println(AbsFunc(v))

    	p := &v
    	fmt.Println(p.Abs())
    	fmt.Println(AbsFunc(*p))
    }
    ```
  In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

<aside>
💡 In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

</aside>

---

## Interfaces

It defines a set of method signatures (names, input parameters, and return types) that a type must implement to satisfy the interface.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

This declares a `Shape` interface with two methods: `Area()` and `Perimeter()`. Any type that wants to be considered a "shape" must implement both of these methods.

```go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.width + 2*r.height
}
// ... similar implementations for Circle, Triangle, etc.
```

```go
func PrintShapeInfo(s Shape) {
    fmt.Println("Area:", s.Area())
    fmt.Println("Perimeter:", s.Perimeter())
}

rect := Rectangle{width: 5, height: 10}
PrintShapeInfo(rect)
```

The `PrintShapeInfo` function accepts any value that implements the `Shape` interface. This allows you to pass it various shapes (rectangles, circles, etc.) without having to write a separate function for each type.

- Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`
  Calling a method on an interface value executes the method of the same name on its underlying type.
  - ```go
    type I interface {
    	M()
    }

    type T struct {
    	S string
    }

    func (t *T) M() {
    	fmt.Println(t.S)
    }

    type F float64

    func (f F) M() {
    	fmt.Println(f)
    }

    func main() {
    	var i I

    	i = &T{"Hello"}
    	describe(i)
    	i.M()

    	i = F(math.Pi)
    	describe(i)
    	i.M()
    }

    func describe(i I) {
    	fmt.Printf("(%v, %T)\n", i, i)
    }
    /*
    (&{Hello}, *main.T)
    Hello
    (3.141592653589793, main.F)
    3.141592653589793
    */
    ```
- Note that an interface value that holds a nil concrete value is itself non-nil. Here `i` holds the value of `t`, which has nil concrete value. So, `t` maybe nil but `i` is not nill.
  ```go
  func main() {
  	var i I

  	var t *T
  	i = t
  	describe(i)
  	i.M()

  	i = &T{"hello"}
  	describe(i)
  	i.M()
  }
  ```
- A nil interface value holds neither value nor concrete type. Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which *concrete* method to call.
  ```go
  func main() {
  	var i I
  	describe(i)
  	i.M()
  }

  func describe(i I) {
  	fmt.Printf("(%v, %T)\n", i, i)
  }
  /*
  (<nil>, <nil>)
  panic: runtime error: invalid memory address or nil pointer dereference
  */
  ```
- An empty interface may hold values of any type. (Every type implements at least zero methods.) Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.
  - ```go
    func main() {
    	var i interface{}
    	describe(i)

    	i = 42
    	describe(i)

    	i = "hello"
    	describe(i)
    }

    func describe(i interface{}) {
    	fmt.Printf("(%v, %T)\n", i, i)
    }
    /*
    (<nil>, <nil>)
    (42, int)
    (hello, string)
    */
    ```

---

### Type Assertion

- A *type assertion* provides access to an interface value's underlying concrete value. `t := i.(T)`
- To *test* whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
  `t, ok := i.(T)`
  If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.
  If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.
  - ```go
    func main() {
    	var i interface{} = "hello"

    	s := i.(string)
    	fmt.Println(s)

    	s, ok := i.(string)
    	fmt.Println(s, ok)

    	f, ok := i.(float64)
    	fmt.Println(f, ok)

    	f = i.(float64) // panic
    	fmt.Println(f)
    }
    /*
    hello
    hello true
    0 false
    panic: interface conversion: interface {} is string, not float64
    */
    ```
- A *type switch* is a construct that permits several type assertions in series. A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
  ```
  switch v := i.(type) {
  case T:
      // here v has type T
  case S:
      // here v has type S
  default:
      // no match; here v has the same type as i
  }
  ```
  - ```go
    func do(i interface{}) {
    	switch v := i.(type) {
    	case int:
    		fmt.Printf("Twice %v is %v\n", v, v*2)
    	case string:
    		fmt.Printf("%q is %v bytes long\n", v, len(v))
    	default:
    		fmt.Printf("I don't know about type %T!\n", v)
    	}
    }

    func main() {
    	do(21)
    	do("hello")
    	do(true)
    }
    /*
    Twice 21 is 42
    "hello" is 5 bytes long
    I don't know about type bool!
    */
    ```

---

### Stringers

The `Stringer` interface is defined in Go's `fmt` package. It's a simple, yet powerful, interface that allows types to define how they should be represented as strings.

```go
type Stringer interface {
    String() string
}
```

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    p := Person{"Alice", 30}
    fmt.Println(p) // Output: Alice (30 years old)
}
```

1. We define a `Person` struct.
2. The `String()` method on `Person` formats the name and age into a readable string.
3. When we print the `Person` value, Go calls the `String()` method, producing the desired output.

- Here is a exercise with `Stringers`
  - ```go
    type IPAddr [4]byte

    // TODO: Add a "String() string" method to IPAddr.
    func ( ip IPAddr) String() string{
    	return fmt.Sprintf("%d.%d.%d.%d",uint8(ip[0]),uint8(ip[1]),uint8(ip[2]),uint8(ip[3]))
    }
    func main() {
    	hosts := map[string]IPAddr{
    		"loopback":  {127, 0, 0, 1},
    		"googleDNS": {8, 8, 8, 8},
    	}
    	fmt.Printf("type of IPAddr: %T\n",hosts["loopback"])
    	for name, ip := range hosts {
    		fmt.Printf("%v: %v\n", name, ip)
    	}
    }
    /*
    type of IPAddr: main.IPAddr
    loopback: 127.0.0.1
    googleDNS: 8.8.8.8
    */
    ```

---

## Errors

The `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

```go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
/*
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
*/
```

Exercise : Errors

- ```go
  type ErrNegativeSqrt float64

  func (e ErrNegativeSqrt) Error() string {
  	return fmt.Sprintf("Cannot Sqrt negative number: %.2f", float64(e))
  }
  func _sqrt(x float64) float64 {
  	const lim = 1e-6
  	sqrtValue := 1.0
  	for itr:=0; itr<10; itr++ {
  		sqrtValue -= ((sqrtValue*sqrtValue - x)/(2*sqrtValue))
  	}
  	return sqrtValue
  }

  func Sqrt(x float64) (float64, error) {
  	if(x>=0) {
  		return _sqrt(x), nil
  	} else {
  		return 0, ErrNegativeSqrt(x)
  	}

  }

  func main() {
  	fmt.Println(Sqrt(2))
  	fmt.Println(Sqrt(-2))
  }
  /*
  1.414213562373095 <nil>
  0 Cannot Sqrt negative number: -2.00
  */
  ```

---

### Readers

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data. The `io.Reader` interface has a `Read` method: `func (T) Read(b []byte) (n int, err error)`

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

```go
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
/*
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
*/
```

Exercise: Readers

- ```go
  type MyReader struct{}

  func (r MyReader) Read(p []byte) (n int, err error) {
      for i := range p {
          p[i] = 'A' // Fill the buffer with 'A' characters
      }
      return len(p), nil // Always read the full buffer
  }

  func main() {
      reader.Validate(MyReader{}) // Test the reader
  }
  /*
  OK!
  */
  ```
