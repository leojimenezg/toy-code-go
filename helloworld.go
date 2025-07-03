package main

// This kind of import is called "factored" import statement.
import (
	"fmt"
	"strings"
	"math/rand"
	"math/cmplx"
	"runtime"
)

// Package level variables
var python, java, rust bool

// Factored variable declaration and different types
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

// Be aware of the argument's structure.
func add(x int, y int) int {
	return x + y
}

// Argument types can be omitted when they are the same, but the last one.
func subtract(x, y int) int {
	return x - y
}

// Multiple return values.
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Naked return, only recommended for short functions.
}

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

// defer functions use a LIFO order stack
func count() int {
	fmt.Println("Counting: ")
	defer fmt.Println("Done!") // First In - Last Out
	for i := range 10 {
		defer fmt.Println(i) // i == 9 Last In - First Out
	}
	return 0 
}

type T struct {
	X, Y, Z int
	a, b, c string
}

type P struct {
	p1, p2 float32
}

func (p P) Sum() float32 {
	return p.p1 + p.p2
}

func (p *P) Scale(f float32) {
	p.p1 = p.p1 * f
	p.p2 = p.p2 * f
}

// Package entry point.
func main() {
	fmt.Println("Random number:", rand.Intn(10))

	fmt.Print("Add of two numbers: ", add(rand.Intn(10), rand.Intn(10)), "\n")
	fmt.Print("Subtract of two numbers: ", subtract(rand.Intn(10), rand.Intn(10)), "\n")

	a, b := swap("Hello", "World")
	fmt.Println("Swaped values:", a, b)

	fmt.Printf("Split values: ")
	fmt.Println(split(17))

	var foo1 int = 15
	foo2 := 18
	fmt.Println(foo1, foo2, python, java, rust)

	var zero int
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(add(zero, int(u)))

	// Normal for loop
	for i:= 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	// Continued foop loop
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	// While (for) loop
	while := 0
	for while < 500 {
		while += 100
	}

	/* Forever loop
	for {
	}
	*/

	var os string = runtime.GOOS // Operating System
	switch os {
	case "darwin":
		fmt.Println("MacOS power")
	case "linux":
		fmt.Println("Linux power")
	case "windows":
		fmt.Println("Fuck windows")
	default:
		fmt.Printf("%s power\n", os)
	}

	count()

	var x int = 42
	var p *int = &x
	fmt.Println("x value from pointer:", *p)

	var v1 T = T{ 1, 2, 3, "a", "b", "c" }
	fmt.Println("Struct v1:", v1)
	var v2 T = T{ Z: 1, Y: 2, X: 3, c: "a", b: "b", a: "c"}
	fmt.Println("Struct v2:", v2)
	var v3 T = T{}
	fmt.Println("Struct v3:", v3)

	var array1 [6]int
	array1[0], array1[1], array1[2] = 1, 2, 3
	array1[3], array1[4], array1[5] = 7, 8, 9
	fmt.Println("Array 1:", array1)

	array2 := [3]int{ 4, 5, 6}
	fmt.Println("Array 2:", array2)

	var slice1 []int = array1[0:3]
	fmt.Println("Slice1 from Array1:", slice1)

	/*
	The difference between a something literal or not is in the use of values:
	something not literal: var numbers [4]int
	something literal: var numbers [4]int = [4]int{1, 2, 3, 4}
	*/

	// This array literal contains all the elements of the array.
	array_literal := [4]bool{true, false, false, true}
	fmt.Println("Array literal:", array_literal)
	// This slice literal creates the array and uses its reference.
	slice_literal := []bool{true, false, false, true}
	fmt.Println("Slice literal:", slice_literal)

	var empty []int
	if empty == nil {
		fmt.Println("Empty slice!")
	}

	// Slice of slices
	board := [][]string {
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"
	fmt.Println("Slice of slices:")
	for i:= 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int = make([]int, 0, 4)
	fmt.Printf("Elements: %v, length: %d, capacity: %d\n", s, len(s), cap(s))
	s = append(s, 0, 1)
	fmt.Printf("Elements: %v, length: %d, capacity: %d\n", s, len(s), cap(s))
	s = append(s, 2, 3)
	fmt.Printf("Elements: %v, length: %d, capacity: %d\n", s, len(s), cap(s))
	s = append(s, 4, 5)
	fmt.Printf("Elements: %v, length: %d, capacity: %d\n", s, len(s), cap(s))
	// When the length of a slice is greater than the slice capacity, a new array generally with doubled the original capacity is created.

	var pow = [8]int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		fmt.Printf("2^%d = %d\n", index, value)
	}

	var map1 map[string]int = make(map[string]int)
	map1["items"] = 3
	map1["price"] = 2
	map1["day"] = 13
	fmt.Println("Map1:", map1)

	// Map literals
	map2 := map[string]string {
		"name": "Mike",
		"lastname": "Thompson",
		"user": "mitho12",
	}
	fmt.Println("Map2:", map2)
	delete(map2, "user")
	fmt.Println("Map2:", map2)

	p1 := P{3, 7}
	fmt.Println("P1 normal:", p1)
	fmt.Println("P1 sum:", p1.Sum())
	p1.Scale(5)
	fmt.Println("P1 scaled:", p1)
}
