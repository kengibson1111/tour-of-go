package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// reslice should be [3, 5, 7]. element for high-bound index (4) not included.
	s = s[1:4]
	fmt.Println(s)

	// reslice of the reslice should be [3, 5]. element for high-bound index (2) not included.
	s = s[:2]
	fmt.Println(s)

	// reslice of the reslice of the reslice should be [5]. high-bound index default is reslice
	// of the reslice length (2). element for high-bound index (2) not included because it does
	// not exist.
	s = s[1:]
	fmt.Println(s)
}
