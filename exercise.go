package hello

import (
	//"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	_, err1 := rot.r.Read(p)
	if err != nil {
		return 0, err1
	}
	for i := range p {
		//fmt.Println(p[i])
		p[i] = p[i] + 13
		//fmt.Println("n",p[i])
	}
	return len(p), nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
