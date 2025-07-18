package main

import "fmt"

func main() {
	if 7%2 == 0 { // here's a basic example.
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // you can have an if statement without an else
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { // logical operator like && and || are often useful in condition.
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { // A statement can preced conditionals; any variables declare in this statement are available in the current and all subsquent branches.
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Note that you don't need parentheses around condition in Go, but that the branches are required'
// tidak ada ternary operator di go, jadi kamu akan membutuhkan if statement bahkan pada basic condition
