package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout,
// grep and sed are common line filters.

// Here an example line filter in go that writes a capitalized version all input text.
// you can use this pattern to write your own go line filters.

func main() {

	// wrapping the unbuffered os.Stdin with a buffered scanner gives us a convinient Scan method that advances the scanner to the next token;
	// which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// text return the current token, here the next line, from the input.
	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl) // write out the uppercased line.
	}

	// Check for errors during Scan.End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
