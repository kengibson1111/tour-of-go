package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fibonaccis []int
	return func() int {
		fiblen := len(fibonaccis)

		switch {
		case fiblen == 0:
			fibonaccis = append(fibonaccis, 0)
		case fiblen == 1:
			fibonaccis = append(fibonaccis, 1)
		default:
			fibonaccis = append(fibonaccis, (fibonaccis[fiblen-1] + fibonaccis[fiblen-2]))
		}

		return fibonaccis[fiblen]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

