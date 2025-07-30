package main

// Reading and writing files are basic tasks needed for many Go programs.
// Frist we'll look at some examples or reading files.

import (
	"fmt"
	"io"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Perhaps the most basic file reading task is slurping a files entire contents into memory.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// you'll often want more control over how and what parts of a file are read.
	// for these task, start by opening a file to obtain an os.File value.'
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the begining of the file. Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also seek to a known location in the file and read from there
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

}
