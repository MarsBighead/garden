package api

import (
	"fmt"
	"testing"
)

func TestGetPb(t *testing.T) {
	data := GetPb()
	fmt.Printf("in test: %v\n", string(data))

}
