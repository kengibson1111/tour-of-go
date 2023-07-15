package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)

	if i == nil {
		fmt.Println("nil interface value")
	} else {
		i.M()
	}
}
