package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this rot13Reader) Read(buff []byte) (n int, err error) {
	n, err = this.r.Read(buff)
	for i := 0; i < n; i++ {
		c := buff[i]
		var off byte
		if c >= 'A' && c <= 'Z' {
			off = 'A'
		}
		if c >= 'a' && c <= 'z' {
			off = 'a'
		}
		if off > 0 {
			buff[i] = (c-off+13)%26 + off
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // prints out «You cracked the code!»
}
