package main

import (
	"fmt"
	"os"
)

// Go offer excellent support for string formatting in the printf tradition.
// Here are some example of common string formatting tasks.

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}               // go offer printing "verbs" designed to format general go values.
	fmt.Printf("struct1: %v\n", p) // for example, this print an instance of our point struct.

	fmt.Printf("struct2: %+v\n", p) // if the value is a struct, the %+v variant will include the struct's field names.

	fmt.Printf("struct3: %#v\n", p) // The %#v variant prints a go syntax representation of the value, i.e. the source code snippet that would produce that value.

	fmt.Printf("type: %T\n", p) // to print the type of a value, use %T.

	fmt.Printf("bool: %t\n", true) // formatting booleans is straight-forward.

	fmt.Printf("int: %d\n", 123) // There are many options for formatting integers. Use %d for standard, base-10 formatting.

	fmt.Printf("bin: %b\n", 14) // this prints a binary representation.

	fmt.Printf("char: %c\n", 33) // This prints the character corresponding to the given integers.

	fmt.Printf("hex: %x\n", 456) // %x provides hex encoding.

	fmt.Printf("float1: %f\n", 78.9) // There are also several formatting options for floats. for basic decimal formatting use %f.

	fmt.Printf("float2: %e\n", 123400000.0) // %e and %E format the float in (slightly different version of) scientific notation.
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"") // for basic string printing use %s.
	fmt.Printf("str2: %q\n", "\"string\"") // to double-quote strings as in go source, use %q
	fmt.Printf("str3: %x\n", "hex this")   // As with integers seen earlier, %x rendes the string in base-16, with two output characters per byte input.

	fmt.Printf("pointer: %p\n", &p) // to print a representation of a pointer, use %p

	fmt.Printf("width1: |%6d|%6d|\n", 12, 234) // when formatting number you will often want to control the width and precision of the resulting figure.
	// to specify the width of an integer, use a number after the % in the verb. By default the result will be the right-justified and padded with the spaces.

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.54)   // you can also specify the width of printed floats, though usually you'll also want to restrict the decimal precision at the same time with the width. precision syntax.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // to left-justify, use the - flag.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")      // you may also want to control width when formatting strings, especially to ensure that they align in table-like output. for basic right-justified width.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")    // to left-justify use - flag as with numbers.

	// So far we've seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and return a string without printing it anywhere.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// you can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
