package closure

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		result := x
		x, y = y, x + y
		return result
	}
}

func main() {
	fib := fibonacci()
	for i:= 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
