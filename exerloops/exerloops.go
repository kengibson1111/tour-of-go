package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 1.0<<64 - 1.0
	for delta >= 0.0001 {
		znew := z - ((z*z - x)/2*z)
		delta = znew - z
		if delta < 0 {
			delta = -delta
		}

		z = znew
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

