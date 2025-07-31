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

	// Other methods of seeking are relative to the current cursor position,
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// and relative to the end of file.
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// The io package provides some functions that may be helpful for file reading.
	// For example, reads like the ones above can be more robustly implemented with ReadAtLeast.

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Print("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// ok reading file saya ulangi sekali lagi besok lah, pusing og.
}
