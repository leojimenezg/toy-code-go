package examples

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	bytesR, errorR := r.reader.Read(b)
	for i := range b[:bytesR] {
		if b[i] > 64 && b[i] < 91 {  // uppercase letters
			for range 13 {
				if b[i] + 1 < 91 {
					b[i] += 1
				} else {
					b[i] = 65
				}
			}
		} else if b[i] > 96 && b[i] < 123 {  // lowercase letters
			for range 13 {
				if b[i] + 1 < 123 {
					b[i] += 1
				} else {
					b[i] = 97
				}
			}
		}
	}
	return bytesR, errorR
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
