package main

// task
// http://127.0.0.1:3999/methods/23

import (
	// "fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n_m int, err_m error) {

	p := make([]byte, 1)
	var result []string
	for {
		n, err := rot.r.Read(p)
		if err == io.EOF {
			rs := strings.Join(result[:], "")
			n_m = copy(b, rs)
			break
		}

		src_rune := []rune(string(p[:n]))[0]
		switch {
		case src_rune >= 'A' && src_rune <= 'Z':
			result = append(result, string('A'+(((src_rune-'A')+13)%26)))
		case src_rune >= 'a' && src_rune <= 'z':
			result = append(result, string('a'+(((src_rune-'a')+13)%26)))
		default:
			result = append(result, string(src_rune))
		}

	}

	return n_m, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.CopyN(os.Stdout, r, 21)
}

