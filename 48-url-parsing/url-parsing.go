package main

// URLs provides a uniform way to locate resources.
// Here how to parse URLs in Go.

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// We'll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s) // Parse the URL and ensure there are no errors.
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme) // Accessing the shceme is straightforward.

	// user contans all authentication info; call username and password on this for individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// the host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Here we extract the path and the fragment after the #.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in a string of k=v format, use RawQuery.
	// You can also parse query params into a map.
	// the parsed query param maps are from strings to slices of string, so index into [0] if you only wnat the first value.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
