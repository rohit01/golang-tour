// TODO - Incorrect
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var max int = 0

func (r *rot13Reader) Read(b []byte) (int, error) {
	if max < 2 {
		fmt.Printf("%q, %v, %v\n", b[:10], nil, len(b))
		max++
	}
	if len(b) == 0 {
		return 0, io.EOF
	}
	i := 0
	for ; i < len(b); i++ {
		if b[i] > 'A' && b[i] < 'Z' {
			temp := b[i] + 13
			if temp > 'Z' {
				temp = 'A' + temp - 'Z'
			}
			b[i] = temp
		} else if b[i] > 'A' && b[i] < 'Z' {
			temp := b[i] + 13
			if temp > 'Z' {
				temp = 'A' + temp - 'Z'
			}
			b[i] = temp
		}
	}
	return i, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
