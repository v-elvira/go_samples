package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

// начало решения

type RandReader struct {
	max int
}

func (r *RandReader) Read(p []byte) (n int, err error) {

	if r.max <= 0 {
		return 0, io.EOF
	}

	if len(p) > r.max {
		p = p[:r.max]
	}

	n, _ = rand.Read(p)
	r.max -= n

	return n, nil
}

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {

	return &RandReader{max: max}
}

// конец решения

func main() {
	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
	// (значения могут отличаться)
}

// // v2: type randomReader struct{}

// func (r *randomReader) Read(p []byte) (n int, err error) {
//     return rand.Read(p)
// }

// func RandomReader(max int) io.Reader {
//     rd := &randomReader{}
//     return io.LimitReader(rd, int64(max))
// }
