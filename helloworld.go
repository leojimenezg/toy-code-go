package main

// This kind of import is called "factored" import statement.
import (
	"fmt"
	"math/rand"
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

// Package entry point.
func main() {
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Print("Add of two numbers: ", add(rand.Intn(10), rand.Intn(10)), "\n")
	fmt.Print("Subtract of two numbers: ", subtract(rand.Intn(10), rand.Intn(10)), "\n")
	a, b := swap("Hello", "World")
	fmt.Println("Swaped values:", a, b)
	fmt.Printf("Split values: ")
	fmt.Println(split(17))
}
