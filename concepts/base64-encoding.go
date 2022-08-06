package concepts

import (
	"encoding/base64"
	"fmt"
)

// Go has built-in functions to encode or decode from base64
func Base64Encoder() {

	s := "ABCDEFG+-xyz"

	// This is how to encode
	strEnc := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(strEnc)
	fmt.Println()

	// This is how to decode
	strDec, _ := base64.StdEncoding.DecodeString(strEnc) // Returns []bytes, and err
	fmt.Println(string(strDec))
	fmt.Println()
}
