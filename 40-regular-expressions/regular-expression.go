package main

import (
	"fmt"
	"regexp"
)

// go offers built-in support for regular expression.
// Here are some example of common regexp-related tasks in go.

func main() {
	// this test wheter a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") // above we used a string pattern directly, but for other regexp tasks you'll need to compile an optimized regexp struct.
	fmt.Println(r.MatchString("priach"))  // true // Many methods are avalable on these struct. Here a match test like we saw earlier.

	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // this also finds the first match but return the start and end indexes for the match instead of the matching text.

	fmt.Println(r.FindStringSubmatch("preach punch")) // the submatch variant include information about both the whole-pattern matches and the submatches within those matches.
	// for example  this will return information for both p([a-z]+)ch and ([a-z]+).
}
