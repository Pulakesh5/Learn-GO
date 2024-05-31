package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	r := bufio.NewReader(rot13.r)
	for index := range p {
		rune, _, err := r.ReadRune()
		if err != nil {
			return 0, err
		}
		if (rune >= 'A' && rune <= 'M') || (rune >= 'a' && rune <= 'm') {
			p[index] = byte(rune + 13)
			fmt.Printf("%c", p[index])
		} else if (rune >= 'N' && rune <= 'Z') || (rune >= 'n' && rune <= 'z') {
			p[index] = byte(rune - 13)
			fmt.Printf("%c", p[index])
		} else {
			p[index] = byte(rune)
			fmt.Printf("%c", p[index])
		}
	}
	return len(p), nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
