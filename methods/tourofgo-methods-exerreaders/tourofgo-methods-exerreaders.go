package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	if b == nil || len(b) <= 0 {
		return 0, nil
	}

	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

