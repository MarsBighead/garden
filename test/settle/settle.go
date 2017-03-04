package settle

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"garden/model/aiqiyi"
	"os"
	"strings"
	"encoding/hex"
	"garden/model/aiqiyi"

	"github.com/golang/protobuf/proto"
)

// EncryptSettle Xiaodu-adx protobuf build
func EncryptSettle() {
	//fmt.Printf("Aiqiyi msg settlement serialization\n")
	price := aesPrice()
	settle := &aiqiyi.Settlement{
		Version: proto.Uint32(717171),
		Price:   price,
	}
	// Marshal data mock aiqiyi settlement
	pData, err := proto.Marshal(settle)
	checkError(err)
	//fmt.Printf("Marshal:\n%v\n", pData)
	settlement := encodeSettlement(string(pData))
	fmt.Printf("settlement encoded by base64: %v\n", settlement)
}

func DecryptSettle() {
	//settlement := "CPPiKxIQQN2I4RWqs0-6lJ37JF6Olw!!"
	settlement := "CPPiKxIQQN2I4RWqs0-6lJ37JF6OlyEh"
	fmt.Printf("37 Original settlement msg|%v\n", settlement)
	data := decodeSettlement(settlement)
	usettle := &aiqiyi.Settlement{}

	err := proto.Unmarshal(data, usettle)
	checkError(err)
	// hex: change data dimensionality
	fmt.Printf("Price  | %v\n", hex.EncodeToString(usettle.Price))
	fmt.Printf("Auth   | %v\n", string(usettle.Auth))
}

//func encryptData(data string) ([]byte, error) {
func decodeSettlement(encrypted string) []byte {
	encoding := aiqiyiEncode()
	fmt.Printf("lenth of encrypted: %v\n", len(encrypted))
	var missing = (4 - len(encrypted)%4) % 4
	encrypted += strings.Repeat("!", missing)
	decoded, err := encoding.DecodeString(encrypted)
	check(err)
	end := len(decoded) - missing
	fmt.Printf("55 decodebase64 is: %v, missing: %v\n", len(decoded), missing)
	fmt.Printf("decodebase64 is: %v\n", decoded[0:end])
	//return decoded[0:end]
	return decoded
}

//func encryptData(data string) ([]byte, error) {
func encodeSettlement(decrypted string) string {
	encoding := aiqiyiEncode()
	var missing = (4 - len(decrypted)%4) % 4
	decrypted += strings.Repeat("!", missing)

	src := []byte(decrypted)
	encoded := encoding.EncodeToString(src)
	fmt.Printf("Length of encoded: %v\n", string(encoded))
	return encoded
}
func aiqiyiEncode() *base64.Encoding {
	dictoinary := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
	encoding := base64.NewEncoding(dictoinary).WithPadding('!')
	return encoding
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func EncryptAESPrice() []byte {
	encodedPrice := "40dd88e115aab34ffa949dfb245e8e97"
	fmt.Printf("encodedPrice is %s.\n", encodedPrice)
	price := "1234567890"
	key, _ := hex.DecodeString("46356afe55fa3cea9cbe73ad442cad47")
	crypted := AesEncrypt(price, key)
	fmt.Printf("ase price: %v\n", string(crypted))
	fmt.Printf("ase price: %v\n", hex.EncodeToString(crypted))
	return crypted
}
func AesEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	// 普通base64编码加密 区别于urlsafe base64
	fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))
	return crypted
}

func AesDecrypt(crypted, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err is:", err)
	}
	blockMode := NewECBDecrypter(block)
	fmt.Printf("Length of crypted: %d\n", len(crypted))
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	fmt.Println("source is :", origData, string(origData))
	return origData
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
