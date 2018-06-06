package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func main() {
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	t1, t2, t3 := tree.New(1), tree.New(1), tree.New(3)
	if Same(t1, t2) {
		fmt.Println("t1 == t2 passed")
	} else {
		fmt.Println("t1 == t2 failed")
	}
	if !Same(t1, t3) {
		fmt.Println("t1 != t3 passed")
	} else {
		fmt.Println("t1 != t3 failed")
	}
}
