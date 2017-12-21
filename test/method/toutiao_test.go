package method

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDecryptToutiao(t *testing.T) {
	finalMsg := "VH2SogAO-2ZUfZKiAA77ZlONOMFBLsR53TOlOg"
	fmt.Printf("Price 13532120, with original finalMsg %v\n", finalMsg)
	price := DecryptToutiao(finalMsg)
	if price == "13532120" {
		fmt.Printf("Price decrypt successfully! Price: %v.\n", price)
	} else {
		fmt.Printf("Price decrypt failed! Price: %v.\n", price)
	}
	value := uint32(13532120)
	// ivTT is []byte iv base64 encode from original finalMsg
	// This part simulate Toutiao encrypt price with the iv from toutiao
	ivTT := "VH2SogAO-2ZUfZKiAA77Zg"
	iv, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(ivTT)
	msg := encyptTotiao(value, iv)
	fmt.Printf("Price %v, with encypted finalMsg %v\n", value, msg)
	if finalMsg == msg {
		fmt.Printf("Simulate toutiao encypt final msg successfully!\n")
	} else {
		fmt.Printf("Simulate toutiao encypt final msg failed!\n")
	}
	iv = []byte(strconv.FormatInt(time.Now().UnixNano(), 10))[0:16]
	value = 1200
	finalMsg = encyptTotiao(value, iv)
	fmt.Printf("Price %v, with encypted finalMsg %v\n", value, finalMsg)
	price = DecryptToutiao(finalMsg)
	if price == "1200" {
		fmt.Printf("Price decrypt successfully! Price: %v.\n", price)
	} else {
		fmt.Printf("Price decrypt failed! Price: %v.\n", price)
	}
}

func encyptTotiao(value uint32, iv []byte) string {
	p := uint64(value)
	iKey := "ebcd1234efgh5678ebcd1234efgh5678"
	eKey := "5678dcba1234abcd5678dcba1234abcd"

	hmacPad := hmac.New(sha1.New, []byte(eKey))
	hmacPad.Write(iv)
	pad := hmacPad.Sum(nil)
	price := make([]byte, 8)
	binary.BigEndian.PutUint64(price, p)
	var encPrice []byte
	for i := range price {
		b := pad[i] ^ price[i]
		encPrice = append(encPrice, b)
	}
	unionPriceIv := make([]byte, 24)
	copy(unionPriceIv, price)
	copy(unionPriceIv[8:], iv)
	hmacSig := hmac.New(sha1.New, []byte(iKey))
	hmacSig.Write(unionPriceIv)
	sig := hmacSig.Sum(nil)
	msg := make([]byte, 28)
	copy(msg, iv)
	copy(msg[16:24], encPrice)
	copy(msg[24:], sig[:4])
	finalMsg := make([]byte, 38)
	base64.URLEncoding.WithPadding(base64.NoPadding).Encode(finalMsg, msg)
	return string(finalMsg)
}
