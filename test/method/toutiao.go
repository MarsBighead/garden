package method

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"log"
	"strconv"
)

// DecryptToutiao Decrypt Jinri Toutiao win price follow
//                Google Real-Time Bidding Protocol:
//                https://developers.google.com/ad-exchange/rtb/response-guide/decrypt-price

func DecryptToutiao(finalMsg string) string {
	iKey := "ebcd1234efgh5678ebcd1234efgh5678"
	eKey := "5678dcba1234abcd5678dcba1234abcd"
	// msg represent the content "iv || enc_price || signature" in
	// "final_message = WebSafeBase64Encode( iv || enc_price || signature )""
	msg, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(finalMsg)
	if err != nil {
		log.Fatal("Jinri Toutiao win notice base64 decode failed!")
	}

	// iv := data[:16]; p := data[16:24]; sig := data[24:]
	iv := msg[:16]
	p := msg[16:24]
	sig := msg[24:]
	hmacPad := hmac.New(sha1.New, []byte(eKey))
	hmacPad.Write(iv)
	pricePad := hmacPad.Sum(nil)
	// enc_price = pad <xor> price
	var price []byte
	for i := range p {
		b := p[i] ^ pricePad[i]
		price = append(price, b)
	}
	hmacConfSig := hmac.New(sha1.New, []byte(iKey))
	// unionPriceIv response content "price || iv" in
	// "conf_sig = hmac(i_key, price || iv)"
	unionPriceIv := make([]byte, 24)
	copy(unionPriceIv, price)
	copy(unionPriceIv[8:], msg[:16])
	hmacConfSig.Write(unionPriceIv)
	confSig := hmacConfSig.Sum(nil)
	success := (string(confSig[:4]) == string(sig))
	if success {
		return strconv.FormatUint(binary.BigEndian.Uint64(price), 10)
	}
	return ""
}
