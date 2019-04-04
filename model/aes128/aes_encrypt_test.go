package aes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptAES(t *testing.T) {
	val := "123457890"
	pass, err := EncryptAES(val)
	assert.NoError(t, err)
	assert.Equal(t, "/urDCY9k/xJeS1rfhOzMwg==", pass, "AES-128 encrypt result error!")
}

func TestDecryptAES(t *testing.T) {
	pass := "/urDCY9k/xJeS1rfhOzMwg=="
	val, err := DecryptAES(pass)
	assert.NoError(t, err)
	assert.Equal(t, "123457890", val, "AES-128 decrypt result error!")

}
