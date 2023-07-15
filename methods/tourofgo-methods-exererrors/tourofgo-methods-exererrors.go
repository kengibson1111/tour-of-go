package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}

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

        return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

