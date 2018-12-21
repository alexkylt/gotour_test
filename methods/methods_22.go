package main

// task
// http://127.0.0.1:3999/methods/22

import "golang.org/x/tour/reader"

// import "fmt"

type MyReader struct {
	name string
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(p []byte) (n int, err error) {
	for {
		n := copy(p, m.name)
		// fmt.Println(n)
		return n, nil
	}
}

func main() {
	reader.Validate(MyReader{"A"})
}

