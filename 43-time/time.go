package main

import (
	"fmt"
	"time"
)

// go offers extensive suport for times and durations;
// here are some examples

func main() {
	p := fmt.Println

	now := time.Now() // We'll start by getting the current time.
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // you can build a time struct by providing the year, month, day, etc. times are alwas associated with location, i.e. time zone.
	p(then)

	// you can extract the various component of the time value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // the monday-sunday weekday is also availabel.

	p(then.Before(now)) //these methods compare two times, testing if the first occurs before, after, or at the same time as the second respectively.
	p(then.After(now))
	p(then.Equal(now))

	// the sub methods returns a duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// we can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// you can use add to advance a time by a given duration, or with a - to move backwards by a duration.
	p(then.Add(diff))
	p(then.Add(-diff))
}
