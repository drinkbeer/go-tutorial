package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
	"runtime"
	"strings"
)

// Execute the code:
// jchen45@Jianbins-MacBook-Pro (master)✗ % go build go.basic.go
// jchen45@Jianbins-MacBook-Pro (master)✗ % ./go.basic

// Install module:
// jchen45@Jianbins-MacBook-Pro (master)✗ % go install time     ~/repo/go-tutorial

var c, python, java bool

func main() {
	basic_go()
	basic_types()
	basic_flow()
	advance_types()
}

func basic_go () {
	fmt.Println("===== Begin: basic_go =====")

	fmt.Println("Hello, world")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	// 55
	fmt.Println(add1(42, 13))

	// 55
	fmt.Println(add2(42, 13))

	// world hello
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 7 10
	fmt.Println(split(17))

	// 0 false false false
	var l int
	fmt.Println(l, c, python, java)

	// 1 2 3 true false no!
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, k, c, python, java)

	fmt.Println("===== End: basic_go =====")
}

var (
	ToBe bool = false
	MaxInt uint64     = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-1 + 12i)
)

// bool (default to be false)
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr  (default to be 0)
// byte // alias for uint8
// rune // alias for int32
//      // represents a Unicode code point
// float32 float64
// complex64 complex128


const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func basic_types() {
	fmt.Println("\n===== Begin: basic_types =====")

	// Print out the types and values of variables
	// Type: bool Value: false
	// Type: uint64 Value: 18446744073709551615
	// Type: complex128 Value: (2.349637693219137+2.553585183501061i)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Default value of variables
	// 0 0 false ""
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversion
	// 3 4 5
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)

	// Constants
	// Hello 世界
	// Happy 3.14 Day
	// Go rules? true
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric Constants are high-precision values.
	// And int can store at maximum a 64-bit integer.
	// 21
	// 0.2
	// 1.2676506002282295e+29
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("===== End: basic_types =====")
}

func basic_flow() {
	fmt.Println()
	fmt.Println("\n===== Begin: basic_flow =====")

	// For loop
	// 45
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The init and post statements in for loop are optional.
	// 1024
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// At that point you can drop the semicolons: C's while is spelled for in Go.
	// 1024
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	fmt.Println(sqrt(2), sqrt(-4))

	// Like for, the if statement can start with a short statement to execute before the condition.
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// Variables declared inside an if short statement are also available inside any of the else blocks.
	fmt.Println(pow2(3, 2, 10), pow2(3, 3, 20))

	// Exercise: Loops and Functions
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	// Switch 1
	// A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// Switch 2
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch 3
	// Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Defer
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")

	// Stacking defers
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")


	fmt.Println("===== End: basic_flow =====")
}

func advance_types() {
	fmt.Println("\n===== Begin: advance_types =====")
	
	// Pointers
	// Go has pointers. A pointer holds the memory address of a value.

	// The type *T is a pointer to a T value. Its zero value is nil.
	// var p *int

	// The & operator generates a pointer to its operand.
	// i := 42
	// p = &i

	// The * operator denotes the pointer's underlying value.
	// fmt.Println(*p) // read i through the pointer p
	// *p = 21         // set i through the pointer p

	// This is known as "dereferencing" or "indirecting".
	// Unlike C, Go has no pointer arithmetic.
	i, j := 42, 1024

	pi := &i         // point to i
	fmt.Println(*pi) // read i through the pointer
	*pi = 21         // set i through the pointer
	fmt.Println(i)   // see the new value of i

	pi = &j         // point to j
	*pi = *pi / 512 // divide j through the pointer
	fmt.Println(j)  // see the new value of j

	// Structs 1
	// A struct is a collection of fields.
	type Vertex struct {
		X int
		Y int
	}
	// type VertexLiterals struct {
	// 	X, Y int
	// }
	fmt.Println(Vertex{1, 2})

	// Struct 2: Fields
	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct 3: Pointers to structs
	// Struct fields can be accessed through a struct pointer.
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However,
	// that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	v = Vertex{1, 2}
	pv := &v
	(*pv).X = 1000
	fmt.Println(v)
	pv.X = 1e9
	fmt.Println(v)

	// Struct 4: Struct Literals
	// A struct literal denotes a newly allocated struct value by listing the values of its fields.
	// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	// The special prefix & returns a pointer to the struct value.
	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p  := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, p, v2, v3)

	// Arrays
	// The type [n]T is an array of n values of type T.
	// The expression
	// var a [10]int
	// declares a variable a as an array of ten integers.
	// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices 1
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view 
	// into the elements of an array. In practice, slices are much more common than arrays.
	// The type []T is a slice with elements of type T.
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	// The following expression creates a slice which includes elements 1 through 3 of a:
	// a[1:4]
	primes = [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices 2: Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a_names := names[0:2]
	b_names := names[1:3]
	fmt.Println(a_names, b_names)

	b_names[0] = "XXX"
	fmt.Println(a_names, b_names)
	fmt.Println(names)

	// Slice 3: literals
	// A slice literal is like an array literal without the length.
	// This is an array literal:
	// [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s_slice_literals := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s_slice_literals)

	// Slice 4: default value
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	// For the array
	// var a [10]int
	// these slice expressions are equivalent:
	/*
	a[0:10]
	a[:10]
	a[0:]
	a[:]
	*/
	s = []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// Slice 5: Slice length and capacity
	/*
	A slice has both a length and a capacity.
	The length of a slice is the number of elements it contains.
	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
	*/
	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Auto extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Slice 6: Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s_nil []int
	fmt.Println(s_nil, len(s_nil), cap(s_nil))
	if s_nil == nil {
		fmt.Println("nil!")
	}

	// Slice 7: Creating a slice with make
	/*
	Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)  // len(a)=5
	To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
	*/
	a_slice_make := make([]int, 5)
	printSliceMake("a", a_slice_make)

	b_slice_make := make([]int, 0, 5)
	printSliceMake("b", b_slice_make)

	c_slice_make := b_slice_make[:2]
	printSliceMake("c", c_slice_make)

	d_slice_make := c_slice_make[2:5]
	printSliceMake("d", d_slice_make)

	// Slice 8: Slices of slices
	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Slice 9: Appending to a slice
	/*
	It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

	func append(s []T, vs ...T) []T
	The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

	The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

	If the backing array of s is too small to fit all the given values, a bigger array will be allocated. The returned slice will point to the newly allocated array.
	*/

	// Understand Go Slices Usage and Internals: https://blog.golang.org/go-slices-usage-and-internals

	// The capacity in this chapter is different from the capacity in the online Tutorial.
	var s_append []int
	printSlice(s_append)

	// append works on nil slices.
	s_append = append(s_append, 0)
	printSlice(s_append)

	// The slice grows as needed.
	s_append = append(s_append, 1)
	printSlice(s_append)

	// We can add more than one element at a time.
	s_append = append(s_append, 2, 3, 4)
	printSlice(s_append)

	// Range 1
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Range continued
	// You can skip the index or value by assigning to _.
	// If you only want the index, drop the , value entirely.
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	// Maps 1
	// A map maps keys to values.
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// The make function returns a map of the given type, initialized and ready for use.
	type Vertex_Map struct {
		Lat, Long float64
	}
	
	var m map[string]Vertex_Map

	m = make(map[string]Vertex_Map)
	m["Bell Labs"] = Vertex_Map{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map 2: Map literals
	// Map literals are like struct literals, but the keys are required.
	var m2 = map[string]Vertex_Map{
		"Bell Labs": Vertex_Map{
			40.68433, -74.39967,
		},
		"Google": Vertex_Map{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// Map 3: Map literals continued
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex_Map{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// Map 4: Mutating Maps
	/*
	Insert or update an element in map m:
	m[key] = elem

	Retrieve an element:
	elem = m[key]

	Delete an element:
	delete(m, key)

	Test that a key is present with a two-value assignment:
	elem, ok = m[key]
	If key is in m, ok is true. If not, ok is false.
	If key is not in the map, then elem is the zero value for the map's element type.
	Note: If elem or ok have not yet been declared you could use a short declaration form:
	elem, ok := m[key]
	*/
	map_mutation := make(map[string]int)

	map_mutation["Answer"] = 42
	fmt.Println("The value:", map_mutation["Answer"])

	map_mutation["Answer"] = 48
	fmt.Println("The value:", map_mutation["Answer"])

	delete(map_mutation, "Answer")
	fmt.Println("The value:", map_mutation["Answer"])

	v_mutation, ok := map_mutation["Answer"]
	fmt.Println("The value:", v_mutation, "Present?", ok)

	// Function values
	// Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function closures
	// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// Exercise: Fibonacci closure
	/*
	Let's have some fun with functions.

	Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
	*/
	// Answer: https://gist.github.com/tetsuok/2281812
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	

	fmt.Println("===== End: advance_types =====")
}

// Two ways of writing function
func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return value. A return without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Exercise: Loops and Functions
func Sqrt(x float64) float64 {
	z := float64(2.0)
	s := float64(0.0)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		if math.Abs(z - s) < 1e-10 {
			break;
		}
		s = z
	}
	return z
}

// Slice 5: Slice length and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Slice 7: Creating a slice with make
func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Exercise: Slices
func Pic(dx, dy int) [][]uint8 {
	/*
		pixel := make([][]uint8, dy)
		data := make([]uint8, dx)
		
		for i := range pixel {
			for j := range data {
				data[j] = uint8((i+j)/2)
			}
			pixel[i] = data
		}
		
		return pixel
	*/
		// Allocate two-dimensioanl array.
		a := make([][]uint8, dy)
		for i := 0; i < dy; i++ {
			a[i] = make([]uint8, dx)
		}
		
		// Do something.
	/*
		for i := 0; i < dy; i++ {
			for j := 0; j < dx; j++ {
				switch {
				case j % 15 == 0:
					a[i][j] = 240
				case j % 3 == 0:
					a[i][j] = 120
				case j % 5 == 0:
					a[i][j] = 150
				default:
					a[i][j] = 100
				}
			}
		}
	*/
	
		for i := range a {
			for j := range a[i] {
				switch {
				case j % 15 == 0:
					a[i][j] = 240
				case j % 3 == 0:
					a[i][j] = 120
				case j % 5 == 0:
					a[i][j] = 150
				default:
					a[i][j] = 100
				}
			}
		}
		return a
	}

// Exercise: Maps
	func WordCount(s string) map[string]int {
		counts := make(map[string]int)
		fields := strings.Fields(s)
		for _, v := range fields {
			counts[v] += 1
		}
		return counts
	}

	// Function values
	func compute(fn func(float64, float64) float64) float64 {
		return fn(3, 4)
	}

	// Function closures
	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	// Exercise: Fibonacci closure
	// fibonacci is a function that returns
	// a function that returns an int.
	func fibonacci() func() int {
		n := 0
		a := 0
		b := 1
		c := a + b
		return func() int {
			var ret int
			switch {
			case n == 0:
				n++
				ret = 0
			case n == 1:
				n++
				ret = 1
			default:
				ret = c
				a = b
				b = c
				c = a + b
			}
			return ret
		}
	}