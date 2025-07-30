package main

// go supports time formating and parsing via pattern-based layouts.

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // here a basic example of formatting a time according RFC3339, using the corresponding layout constant.

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1, e)

	/*
		Format and Parse use example-based layouts.
		Usually you'll use a constant from time for these layouts, but you can also supply custom layouts.
		Layouts must use reference time Mon Jan 2 15:14:05 MST 2006 to show the pattern with which to format/parse a given time/string.
		Then example time must be exactly as shown: the year 2006, 15 for the hour, Monday for day of the week, etc.
	*/
	p(t.Format("3:02PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("Mon 2 Jan 2007")) // hasil modifikasi sendiri
	p(t.Format("2006-01-02"))     // hasil modifikasi sendiri

	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2, e)

	// for purely numeric representations you can also use standard string formatting with the extracted components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00-00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
