package bt

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// In-order walk.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// Assumes the trees' sizes can be different.
func Same1(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	same := true
	// Gets all the values from both channels
	// to avoid goroutine leaks.
	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2
		if same && (ok1 != ok2 || v1 != v2) {
			same = false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return same
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// Assumes the trees' sizes are the same.
func Same2(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for v1 := range ch1 {
		v2 := <- ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	same := Same1(tree.New(1), tree.New(1))
	fmt.Println(same)
}
