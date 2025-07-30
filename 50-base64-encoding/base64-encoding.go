package main

/*
	go provides build-in support for base64-encoding/decoding.
*/

import (
	b64 "encoding/base64" // this syntax imports the encoding/base64 package with the b64 name instead of the default base64. It'll save us some space below.
	"fmt"
)

func main() {
	data := "abc123!?$*&()`-=~" // Here the string we'll encode/decode.

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // Go supports both standard and URL-compatible base64. Here how to encode using the standard encoder.
	fmt.Println(sEnc)                                    // the encoder requires a []byte so we can convert our string to that type.

	sDec, _ := b64.StdEncoding.DecodeString(sEnc) // Decoding may return an error, which you can check if you don't already know the input to be well-formerd.
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) // this encodes/decodes using a URL-compatible base64 format.
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
