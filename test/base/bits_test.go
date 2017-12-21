package base

import (
	"fmt"
	"io"
	"testing"
)

func TestReaderToBytes(t *testing.T) {
	const size = 18
	robert := &StringPair{"Robert L.", "Stevenson"}
	david := StringPair{"Davide ", "Zhang"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := readerToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Raw %q\n", raw)
	}
}
func TestReadBytes(t *testing.T) {
	ReadBytes()
}
