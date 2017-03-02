package settle

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDecryptSettle(t *testing.T) {

	price := EncryptAESPrice()
	fmt.Printf("Encrypt price: %v\n", string(price))
	fmt.Printf("Encrypt price(hex): %v\n", hex.EncodeToString(price))
	// EncodeSettle(price)
	hexPrice := DecodeSettle()
	fmt.Printf("Decode price(hex): %v\n", hexPrice)
}
