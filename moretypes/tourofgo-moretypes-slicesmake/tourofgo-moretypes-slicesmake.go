package main

import "fmt"

func main() {
	a := make([]int, 5)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	printSlice("a", a)

	b := a[:0]
	printSlice("b", b)

	c := a[1:3]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
