package util

import (
	"fmt"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	val := "123457890"
	param := EncryptAES(val)
	fmt.Printf("AES encrypt: %v\n", param)
}

func TestDecryptAES(t *testing.T) {
	src := "/urDCY9k/xJeS1rfhOzMwg=="
	value, _ := DecryptAES(src)
	fmt.Printf("AES Decrypt: %v\n", value)
}
