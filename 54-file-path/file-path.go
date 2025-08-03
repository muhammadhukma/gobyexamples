package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// hte filepath package provides functions to parse and construct file path in a way that is portable between operating system;
// dir/file on linuz vs dir\file on windows, for example.

func main() {
	// join should be used to construct path in a portable way.
	// it takes any number of arguments and constructs a hierarchical path from them.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// You should always use Join insted of concatenating / s or \s manually.
	// In addition to providing portability, Join will aslo normalize path by removing sperfluous separatros and deirectory changes,
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to split a path to the directory and the file.
	// Alternatively, Split will return both in the same call.
	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// some file names have extensions following a dot.
	// We can split the extension out of such names with Ext (Extension).
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// To find the file's name with extension removed, use strings.TrimSuffix.

	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target.
	// It returns an error if the target cannot be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
