package concepts

import (
	"crypto/sha256"
	"fmt"
)

// Go has built-in functions for encoding using sha256
// Here is how we can hash strings and output them as bytes or as a hex
func Hashing() {
	str := "string to be hashed"

	hasher := sha256.New()

	hasher.Write([]byte(str))

	byteSlice := hasher.Sum(nil)

	fmt.Println(str)

	// This will print it as a readable hash
	fmt.Printf("%x\n", byteSlice)

	// This will print it as bytes
	fmt.Println(byteSlice)
}
