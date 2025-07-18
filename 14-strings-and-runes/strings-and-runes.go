package main

import (
	"fmt"
	"unicode/utf8"
)

// a go strings is a read-only slice of bytes.
// the language and the standard library treat strings specially - as containers of text endcoded in UTF-8.
// In other languages, strings are made of "characters".
// in go, the concept of character is called a rune - it's an integer that represent a unicode code point.

func main() {
	// s is a string assigned a literal value representing the word "hello" in the Thai language.
	// go string literals are UTF-8 endcoded text
	const s = "สวัสดี"

	// since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	// indexing into a string produces the raw byte values at each index. this loop generates the hex value of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// to count how many runes are in a string, we can use the utf8 package.
	// note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially.
	// some thai characters are represent by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
	for idx, runeValue := range s {
		fmt.Printf("%#U start at %d\n", runeValue, idx)
	}

	// we can achieve the same iteration by using the utf8. DecodeRuneInString function explicitly
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// this demonstrates passing a rune value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")

	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
