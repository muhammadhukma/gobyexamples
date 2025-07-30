package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for binary or text blobs.
// For example, TLS/SSL certificates use SHA256 to compute a certificate's signature.
// Here's ho to compute SHA256 hashes in Go.

func main() {
	s := "sha256 this string"

	h := sha256.New()  // Here we start a new hash.
	h.Write([]byte(s)) // Write expects bytes. If you having a string s, use []bytes(s) to coerce it to bytes.

	bs := h.Sum(nil) // This gets the finalized hash result as a byte slice.
	// the argument to Sum can be used to append to an existing byte slice: it usually isn't needed.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
