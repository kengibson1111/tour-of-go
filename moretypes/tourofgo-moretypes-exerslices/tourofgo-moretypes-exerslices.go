package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pixels [][]uint8

	for x := 0; x < dx; x++ {
		columns := make([]uint8, dy)
		for y := 0; y < dy; y++ {
			columns[y] = uint8(((x*y)/2) % (1<<8-1))
		}

		pixels = append(pixels, columns)
	}

	return pixels
}

func main() {
	pic.Show(Pic)
}

