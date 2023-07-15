package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ints1 := make([]int, 5)
	for i := range ch1 {
		fmt.Println(i)
		ints1 = append(ints1, i)
	}

	ch2 := make(chan int)
	go Walk(t2, ch2)

	ints2 := make([]int, 5)
	for i := range ch2 {
		fmt.Println(i)
		ints2 = append(ints2, i)
	}

	if len(ints1) != len(ints2) {
		return false
	}

	for i := 0; i < len(ints1); i++ {
		if ints1[i] != ints2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

