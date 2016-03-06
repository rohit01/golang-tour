package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for i = 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
