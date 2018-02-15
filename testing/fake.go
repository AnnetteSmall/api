package main

import (
	//"fmt"
	"io"
	"fmt"
)

type FakeReader struct {
}

func (FakeReader) Read(p []byte) (n int, err error) {
	// return an integer and error or nil
	n = 5
	err = nil
	return n, err
}

func ReadAllTheBytes(reader io.Reader) []byte {
	// read from the reader..

}

func main() {
	fakeReader := FakeReader{}
	// You could create a method called SetFakeBytes which initialises canned data.
	//fakeReader.SetFakeBytes([]byte("when called, return this data"))
	bytes := ReadAllTheBytes(fakeReader)
	fmt.Printf("%d bytes read.\n", len(bytes))
}
