package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"log"
)

//EncryptAES  AES-128  length of key: 16, 24, 32 bytes is mapping AES-128, AES-192, AES-256
func EncryptAES(src string) (based string, err error) {
	//Initial []byte token
	token, err := hex.DecodeString("46356afe55fa3cea9cbe73ad442cad47")
	if err != nil {
		log.Fatal(err)
	}
	// Block from Cipher
	block, err := aes.NewCipher(token)
	if err != nil {
		log.Fatal(err)
		return
	}
	blockSize := block.BlockSize()
	ecbe := cipher.NewCBCEncrypter(block, token[:blockSize])
	content := PKCS5Padding([]byte(src), blockSize)
	// Initial crypt value
	crypted := make([]byte, len(content))
	ecbe.CryptBlocks(crypted, content)
	based = base64.StdEncoding.EncodeToString(crypted)
	return
}

// DecryptAES Decrypt AES encrypted value
func DecryptAES(src string) (value string, err error) {
	buf, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := hex.DecodeString("46356afe55fa3cea9cbe73ad442cad47")
	if err != nil {
		log.Fatal(err)
		return
	}
	block, err := aes.NewCipher(token)
	if err != nil {
		log.Fatal(err)
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, token[:blockSize])
	byteValue := make([]byte, len(buf))
	blockMode.CryptBlocks(byteValue, buf)
	value = string(PKCS5UnPadding(byteValue))
	return
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// Remove the last unpadding times
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
